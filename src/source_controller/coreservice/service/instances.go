/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.,
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the ",License",); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an ",AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/http/rest"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/source_controller/coreservice/bussiness/cache"
	"configcenter/src/source_controller/coreservice/multilingual"
)

func (s *coreService) CreateOneModelInstance(ctx *rest.Contexts) {
	inputData := metadata.CreateModelInstance{}
	if err := ctx.DecodeInto(&inputData); nil != err {
		ctx.RespAutoError(err)
		return
	}
	objectID := ctx.Request.PathParameter("bk_obj_id")
	dataResult, err := s.core.InstanceOperation().CreateModelInstance(ctx.Kit, objectID, inputData)
	if err != nil {
		blog.Errorf("CreateOneModelInstances err:%v", err)
		ctx.RespEntityWithError(dataResult, err)
		return
	}
	ids := make([]uint64, 0, 1)
	ids = append(ids, dataResult.Created.ID)

	sc := &cache.SendCache{
		Cond:     mapstr.MapStr{"bk_inst_id": mapstr.MapStr{common.BKDBIN: ids}},
		Core:     s.core,
		ObjectID: objectID,
	}

	sc.CopyKit(ctx.Kit)

	if err == nil {
		cache.SendCacheChan <- sc
		//go func() {

		//var iscache bool
		//var ui interface{}
		//ui, iscache = cache.CacheObjMap.Get(objectID)
		//if !iscache {
		//	blog.Errorf("objid:%s not found cache skip", objectID)
		//	return
		//}
		//ids := make([]uint64, 0, 1)
		//ids = append(ids, dataResult.Created.ID)
		//for _, create := range dataResult.CreateManyInfoResult.Created {
		//	ids = append(ids, create.ID)
		//}
		//cond := mapstr.MapStr{"bk_inst_id": mapstr.MapStr{common.BKDBIN: ids}}
		//err, _ := sc.SearchDBAndSaveCache(objectID, ui.(string), s.core, iscache)
		//if err != nil {
		//	blog.Errorf("objid:%s unique:%v save redis err:%v", objectID, ui, err)
		//	return
		//}

		//}()
	}
	ctx.RespEntityWithError(dataResult, err)
}

func (s *coreService) InsertManyModelInstances(ctx *rest.Contexts) {
	inputData := metadata.CreateManyModelInstance{}
	if err := ctx.DecodeInto(&inputData); nil != err {
		ctx.RespAutoError(err)
		return
	}
	objectID := ctx.Request.PathParameter("bk_obj_id")
	//fmt.Println("11111111111111114444444444 InsertManyModelInstances", ctx.Kit)
	ids, err := s.core.InstanceOperation().InsertManyModelInstance(ctx.Kit, objectID, inputData)
	if err != nil {
		blog.Errorf("InsertManyModelInstances err:%v", err)
		ctx.RespEntityWithError(nil, err)
		return
	}
	//
	//ids := make([]uint64, 0, len(dataResult.Created))
	//for _, create := range dataResult.CreateManyInfoResult.Created {
	//	ids = append(ids, create.ID)
	//}
	sc := &cache.SendCache{
		Cond:     mapstr.MapStr{"bk_inst_id": mapstr.MapStr{common.BKDBGTE: ids[0], common.BKDBLT: ids[len(ids)-1]}},
		Core:     s.core,
		ObjectID: objectID,
	}

	sc.CopyKit(ctx.Kit)

	if err == nil {

		cache.SendCacheChan <- sc

	}
	ctx.RespEntityWithError(map[string]interface{}{"success": len(ids)}, err)
}

func (s *coreService) UpdateManyModelInstances(ctx *rest.Contexts) {
	inputData := metadata.CreateManyModelInstance{}
	if err := ctx.DecodeInto(&inputData); nil != err {
		ctx.RespAutoError(err)
		return
	}
	objectID := ctx.Request.PathParameter("bk_obj_id")
	dataResult, err := s.core.InstanceOperation().UpdateManyModelInstance(ctx.Kit, objectID, inputData)
	if err != nil {
		blog.Errorf("UpdateManyModelInstances err:%v", err)
		ctx.RespEntityWithError(dataResult, err)
		return
	}

	sc := &cache.SendCache{
		Core:     s.core,
		ObjectID: objectID,
	}
	sc.IsCache()

	if sc.Iscache {
		//inputCopy := inputata.Clone()
		datas := &metadata.QueryResult{
			Count: uint64(len(inputData.Datas)),
			Info:  inputData.Datas,
		}
		err, _ := sc.DeleteCache(datas)
		//fmt.Println(dataResult)
		//fmt.Println(err)
		if err != nil {
			blog.Errorf("UpdateManyModelInstances err:%v", err)

		}

	}
	ctx.RespEntityWithError(dataResult, err)
}

func (s *coreService) CreateManyModelInstances(ctx *rest.Contexts) {
	//fmt.Println(ctx.Request.Request.Method)
	inputData := metadata.CreateManyModelInstance{}
	if err := ctx.DecodeInto(&inputData); nil != err {
		ctx.RespAutoError(err)
		return
	}
	objectID := ctx.Request.PathParameter("bk_obj_id")
	dataResult, err := s.core.InstanceOperation().CreateManyModelInstance(ctx.Kit, objectID, inputData)
	if err != nil {
		blog.Errorf("CreateManyModelInstances err:%v", err)
		ctx.RespEntityWithError(dataResult, err)
		return
	}

	ids := make([]uint64, 0, len(dataResult.Created))
	for _, create := range dataResult.CreateManyInfoResult.Created {
		ids = append(ids, create.ID)
	}
	sc := &cache.SendCache{
		Cond:     mapstr.MapStr{"bk_inst_id": mapstr.MapStr{common.BKDBIN: ids}},
		Core:     s.core,
		ObjectID: objectID,
	}

	sc.CopyKit(ctx.Kit)

	if err == nil {

		cache.SendCacheChan <- sc

	}
	ctx.RespEntityWithError(dataResult, err)
}

func (s *coreService) UpdateModelInstances(ctx *rest.Contexts) {
	inputData := metadata.UpdateOption{}
	if err := ctx.DecodeInto(&inputData); nil != err {
		ctx.RespAutoError(err)
		return
	}
	objectID := ctx.Request.PathParameter("bk_obj_id")
	dataResult, err := s.core.InstanceOperation().UpdateModelInstance(ctx.Kit, objectID, inputData)
	if err != nil {
		blog.Errorf(" UpdateModelInstances err:%v", err)
		ctx.RespEntityWithError(dataResult, err)
		return
	}

	//var iscache bool
	//var ui interface{}
	//var unique string
	//ui, iscache = cache.CacheObjMap.Get(objectID)
	//unique = ui.(string)
	sc := &cache.SendCache{
		Cond:     inputData.Condition,
		Core:     s.core,
		ObjectID: objectID,
	}
	sc.IsCache()
	if sc.Iscache {

		sc.CopyKit(ctx.Kit)
		//inputCopy := inputata.Clone()
		err, _ := sc.SearchDBAndSaveCache()
		//fmt.Println(dataResult)
		//fmt.Println(err)
		if err != nil {
			blog.Errorf("Update Instance Cache save Cache Cond:[%v] err:%v", inputData.Condition, err)

		}

	}

	ctx.RespEntityWithError(dataResult, err)
}

func (s *coreService) SearchModelInstances(ctx *rest.Contexts) {

	//dd := s.core.CacheOperation().GetMany(context.Background(), []string{"cc:services:endpoints:apiserver"})
	//dd := s.engine.RedisClient.GetMany(context.Background(), []string{"cc:services:endpoints:apiserver"})
	//fmt.Println(dd)
	inputData := metadata.QueryCondition{}
	if err := ctx.DecodeInto(&inputData); nil != err {
		ctx.RespAutoError(err)
		return
	}

	objectID := ctx.Request.PathParameter("bk_obj_id")

	// 判断是否有要根据default字段，需要国际化的内容
	if _, ok := multilingual.BuildInInstanceNamePkg[objectID]; ok {
		// 大于两个字段
		if len(inputData.Fields) > 1 {
			inputData.Fields = append(inputData.Fields, common.BKDefaultField)
		} else if len(inputData.Fields) == 1 && inputData.Fields[0] != "" {
			// 只有一个字段，如果字段为空白字符，则不处理
			inputData.Fields = append(inputData.Fields, common.BKDefaultField)
		}
	}

	dataResult, err := s.core.InstanceOperation().SearchModelInstance(ctx.Kit, objectID, inputData)
	if nil != err {
		ctx.RespEntityWithError(dataResult, err)
		return
	}

	multilingual.TranslateInstanceName(s.Language(ctx.Kit.Header), objectID, dataResult.Info)

	ctx.RespEntity(dataResult)
}

func (s *coreService) DeleteModelInstances(ctx *rest.Contexts) {
	inputData := metadata.DeleteOption{}
	if err := ctx.DecodeInto(&inputData); nil != err {
		ctx.RespAutoError(err)
		return
	}
	objectID := ctx.Request.PathParameter("bk_obj_id")

	inputMongoData := metadata.QueryCondition{Condition: inputData.Condition}

	dataResult, err := s.core.InstanceOperation().SearchModelInstance(ctx.Kit, objectID, inputMongoData)
	if nil != err {
		ctx.RespEntityWithError(dataResult, err)
	}

	delResult, err := s.core.InstanceOperation().DeleteModelInstance(ctx.Kit, objectID, inputData)
	if err == nil {
		//var iscache bool
		//var ui interface{}
		//var unique string
		//ui, iscache = cache.CacheObjMap.Get(objectID)
		//unique = ui.(string)
		sc := &cache.SendCache{
			Cond:     inputData.Condition,
			Core:     s.core,
			ObjectID: objectID,
		}
		sc.IsCache()

		if sc.Iscache {

			//inputCopy := inputata.Clone()
			err, _ := sc.DeleteCache(dataResult)
			//fmt.Println(dataResult)
			//fmt.Println(err)
			if err != nil {
				blog.Errorf("DeleteSkipAModelInstances Cond:[%v] err:%v", inputData.Condition, err)

			}

		}
	}
	ctx.RespEntityWithError(delResult, err)
}

func (s *coreService) DeleteSkipArchiveModelInstances(ctx *rest.Contexts) {
	inputData := metadata.DeleteOption{}
	if err := ctx.DecodeInto(&inputData); nil != err {
		ctx.RespAutoError(err)
		return
	}
	objectID := ctx.Request.PathParameter("bk_obj_id")

	inputMongoData := metadata.QueryCondition{Condition: inputData.Condition}

	dataResult, err := s.core.InstanceOperation().SearchModelInstance(ctx.Kit, objectID, inputMongoData)
	if nil != err {
		ctx.RespEntityWithError(dataResult, err)
	}

	err = s.core.InstanceOperation().DeleteArchiveModelInstance(ctx.Kit, objectID, inputData)
	if err == nil {

		sc := &cache.SendCache{
			Cond:     inputData.Condition,
			Core:     s.core,
			ObjectID: objectID,
		}
		sc.IsCache()

		if sc.Iscache {

			//inputCopy := inputata.Clone()
			err, _ := sc.DeleteCache(dataResult)
			//fmt.Println(dataResult)
			//fmt.Println(err)
			if err != nil {
				blog.Errorf("DeleteModelInstances Cond:[%v] err:%v", inputData.Condition, err)

			}

		}
	}
	ctx.RespEntityWithError(nil, err)
}

func (s *coreService) CascadeDeleteModelInstances(ctx *rest.Contexts) {
	inputData := metadata.DeleteOption{}
	if err := ctx.DecodeInto(&inputData); nil != err {
		ctx.RespAutoError(err)
		return
	}
	ctx.RespEntityWithError(s.core.InstanceOperation().CascadeDeleteModelInstance(ctx.Kit, ctx.Request.PathParameter("bk_obj_id"), inputData))
}
