package cache

import (
	"configcenter/src/common/blog"
	"configcenter/src/storage/dal/redis"
	"context"
	"fmt"
	redisv7 "github.com/go-redis/redis/v7"
	"strconv"
	"strings"
	"time"
)

/**
 * @Author: Tao Jun
 * @Since: 2022/7/6
 * @Desc: syncunique.go
**/

var RedisCli *redisv7.Client
var RedisInter redis.Client

var CacheDataChan = make(chan *CacheData, 2048)

func initRedis(cfg redis.Config) error {
	dbNum, err := strconv.Atoi(cfg.Database)
	if nil != err {
		return err
	}
	if cfg.MaxOpenConns == 0 {
		cfg.MaxOpenConns = 3000
	}

	//var client Client
	if cfg.MasterName == "" {
		option := &redisv7.Options{
			Addr:     cfg.Address,
			Password: cfg.Password,
			DB:       dbNum,
			PoolSize: cfg.MaxOpenConns,
		}
		RedisCli = redisv7.NewClient(option)
	} else {
		hosts := strings.Split(cfg.Address, ",")
		option := &redisv7.FailoverOptions{
			MasterName:       cfg.MasterName,
			SentinelAddrs:    hosts,
			Password:         cfg.Password,
			DB:               dbNum,
			PoolSize:         cfg.MaxOpenConns,
			SentinelPassword: cfg.SentinelPassword,
		}
		RedisCli = redisv7.NewFailoverClient(option)
	}

	err = RedisCli.Ping().Err()
	if err != nil {
		return err
	}

	return err
}

func InitCli(ctx context.Context, rediscfg redis.Config, redisInter redis.Client) {
	topCtx = ctx
	RedisInter = redisInter
	for {
		// util redisCli success
		err := initRedis(rediscfg)
		if err == nil {
			break
		} else {
			blog.Errorf("redisCli init fail:%v %d sec after try", err)
		}
		time.Sleep(time.Second * 30)
	}
}
func SyncCacheData() {

	go func() {
		redisCrudChan()
	}()
}

func redisCrudChan() {
	for {
		select {
		case data := <-CacheDataChan:
			if data.Type == Update {
				go func() {
					fmt.Println("update redis chan")
					//blog.Errorf("update redis chan obj:%s unique:%s ", data.ObjID, data.unique)
					timeoutCtx, _ := context.WithTimeout(topCtx, time.Duration(10*time.Second))
					//json, err := data.Data.ToJSON()
					//if err != nil {
					//	blog.Errorf("update redis chan obj:%s unique:%s to json err:%v", data.ObjID, data.Unique, err)
					//}
					err := RedisInter.SetMany(timeoutCtx, data.CacheSaveData).Err()
					if err != nil {
						blog.Errorf("update redis chan obj:%s unique:%s err:%v", data.ObjID, data.Unique, err)
					} else {
						blog.Infof("update redis chan obj:%s unique:%s ", data.ObjID, data.Unique)
					}
				}()
			}
			if data.Type == Delete {
				go func() {
					fmt.Println("delete redis chan")
					timeoutCtx, _ := context.WithTimeout(topCtx, time.Duration(10*time.Second))
					//json, err := data.Data.ToJSON()
					//if err != nil {
					//	blog.Errorf("update redis chan obj:%s unique:%s to json err:%v", data.ObjID, data.Unique, err)
					//}
					err := RedisInter.DelMany(timeoutCtx, data.CacheDelKeys).Err()
					if err != nil {
						blog.Errorf("delete redis chan obj:%s unique:%s err:%v", data.ObjID, data.Unique, err)
					} else {
						blog.Infof("delete redis chan obj:%s unique:%s ", data.ObjID, data.Unique)
					}
				}()
			}
		default:
			continue
		}
	}
}
