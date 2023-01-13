package cache

import (
	"configcenter/src/common/blog"
	"context"
	"fmt"
	"time"
)

/**
 * @Author: Tao Jun
 * @Since: 2022/7/6
 * @Desc: syncunique.go
**/

var CacheDataChan = make(chan *CacheData, 2048)

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
					RedisInter.SetManyExpiration(topCtx, data.CacheSaveData, time.Duration(24*time.Hour))
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
		}
	}
}
