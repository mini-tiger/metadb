/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package app

import (
	"context"
	"fmt"
	"time"

	"configcenter/src/ac/iam"
	"configcenter/src/common/auth"
	"configcenter/src/common/backbone"
	cc "configcenter/src/common/backbone/configcenter"
	"configcenter/src/common/blog"
	"configcenter/src/common/resource/esb"
	"configcenter/src/common/types"
	"configcenter/src/scene_server/admin_server/app/options"
	"configcenter/src/scene_server/admin_server/configures"
	svc "configcenter/src/scene_server/admin_server/service"
	"configcenter/src/storage/dal/mongo/local"
	"configcenter/src/storage/dal/redis"
)

func Run(ctx context.Context, cancel context.CancelFunc, op *options.ServerOption) error {
	// init esb client
	esb.InitEsbClient(nil)

	svrInfo, err := types.NewServerInfo(op.ServConf)
	if err != nil {
		return fmt.Errorf("wrap server info failed, err: %v", err)
	}

	process := new(MigrateServer)
	process.Config = new(options.Config)
	//if os.Getenv("dev") == "true"{
	//
	//}
	// xxx 读取 migrate.yaml 配置文件
	if err := cc.SetMigrateFromFile(op.ServConf.ExConfig); err != nil {
		return fmt.Errorf("parse config file error %s", err.Error())
	}

	process.Config.Errors.Res, _ = cc.String("errors.res")
	process.Config.Language.Res, _ = cc.String("language.res")
	process.Config.Configures.Dir, _ = cc.String("confs.dir")
	process.Config.Register.Address, _ = cc.String("registerServer.addrs")
	snapDataID, _ := cc.Int("hostsnap.dataID")
	process.Config.SnapDataID = int64(snapDataID)

	// load mongodb, redis and common config from configure directory
	mongodbPath := process.Config.Configures.Dir + "/" + types.CCConfigureMongo
	if err := cc.SetMongodbFromFile(mongodbPath); err != nil {
		return fmt.Errorf("parse mongodb config from file[%s] failed, err: %v", mongodbPath, err)
	}

	redisPath := process.Config.Configures.Dir + "/" + types.CCConfigureRedis
	if err := cc.SetRedisFromFile(redisPath); err != nil {
		return fmt.Errorf("parse redis config from file[%s] failed, err: %v", redisPath, err)
	}

	commonPath := process.Config.Configures.Dir + "/" + types.CCConfigureCommon
	if err := cc.SetCommonFromFile(commonPath); err != nil {
		return fmt.Errorf("parse common config from file[%s] failed, err: %v", commonPath, err)
	}

	mongoConf, err := cc.Mongo("mongodb")
	if err != nil {
		return err
	}
	process.Config.MongoDB = mongoConf

	watchDBConf, err := cc.Mongo("watch")
	if err != nil {
		return err
	}
	process.Config.WatchDB = watchDBConf

	redisConf, err := cc.Redis("redis")
	if err != nil {
		return err
	}
	process.Config.Redis = redisConf

	snapRedisConf, err := cc.Redis("redis.snap")
	if err != nil {
		return fmt.Errorf("get host snapshot redis configuration failed, err: %v", err)
	}
	process.Config.SnapRedis = snapRedisConf

	process.Config.Iam, err = iam.ParseConfigFromKV("authServer", nil)
	if err != nil && auth.EnableAuthorize() {
		blog.Errorf("parse iam error: %v", err)
	}

	input := &backbone.BackboneParameter{
		ConfigUpdate: process.onHostConfigUpdate,
		ConfigPath:   op.ServConf.ExConfig,
		Regdiscv:     process.Config.Register.Address,
		SrvInfo:      svrInfo,
	}
	engine, err := backbone.NewBackbone(ctx, input, redisConf)
	if err != nil {
		return fmt.Errorf("new backbone failed, err: %v", err)
	}

	process.Core = engine
	process.ConfigCenter = configures.NewConfCenter(ctx, engine.ServiceManageClient(), redisConf)
	//process.ConfigCenter = configures.NewConfCenter(ctx, engine.ServiceManageClient())

	// adminserver conf not depend discovery
	//xxx redis mongo common 配置写入redis
	err = process.ConfigCenter.Start(
		process.Config.Configures.Dir,
		process.Config.Errors.Res,
		process.Config.Language.Res,
	)

	//xxx fix  配置文件修改 更新redis
	cc.SetFileChange(func() {
		if err := process.ConfigCenter.WriteConfs2Center(process.Config.Configures.Dir); err != nil {
			blog.Errorf("fail to write configures to center, err:%s", err.Error())

		} else {
			blog.Infof("write all configures resource to center %v success", types.CC_SERVCONF_BASEPATH)
		}
	})

	if err != nil {
		return err
	}

	service := svc.NewService(ctx)
	service.Engine = engine
	service.Config = *process.Config
	service.ConfigCenter = process.ConfigCenter
	process.Service = service

	for {
		if process.Config == nil {
			time.Sleep(time.Second * 2)
			blog.V(3).Info("config not found, retry 2s later")
			continue
		}

		db, err := local.NewMgo(process.Config.MongoDB.GetMongoConf(), time.Minute)
		if err != nil {
			return fmt.Errorf("connect mongo server failed %s", err.Error())
		}
		process.Service.SetDB(db)

		watchDB, err := local.NewMgo(process.Config.WatchDB.GetMongoConf(), time.Minute)
		if err != nil {
			return fmt.Errorf("connect watch mongo server failed, err: %v", err)
		}
		process.Service.SetWatchDB(watchDB)

		cache, err := redis.NewFromConfig(process.Config.Redis)
		if err != nil {
			return fmt.Errorf("connect redis server failed, err: %s", err.Error())
		}
		process.Service.SetCache(cache)

		if auth.EnableAuthorize() {
			blog.Info("enable auth center access.")

			iamCli, err := iam.NewIam(nil, process.Config.Iam, engine.Metric().Registry())
			if err != nil {
				return fmt.Errorf("new iam client failed: %v", err)
			}
			process.Service.SetIam(iamCli)
		} else {
			blog.Infof("disable auth center access.")
		}

		if esbConfig, err := esb.ParseEsbConfig(""); err == nil {
			esb.UpdateEsbConfig(*esbConfig)
		}
		break
	}
	err = backbone.StartServer(ctx, cancel, engine, service.WebService(), true)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
	}
	blog.V(0).Info("process stopped")
	return nil
}

type MigrateServer struct {
	Core         *backbone.Engine
	Config       *options.Config
	Service      *svc.Service
	ConfigCenter *configures.ConfCenter
}

func (h *MigrateServer) onHostConfigUpdate(previous, current cc.ProcessConfig) {}
