package cache

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/json"
	"configcenter/src/source_controller/coreservice/core/operation"
	"configcenter/src/storage/driver/mongodb"
	"context"
	"fmt"
	"time"
)

/**
 * @Author: Tao Jun
 * @Since: 2022/7/6
 * @Desc: syncunique.go
**/

func ObjAllDataInitCache() {

	// wait CacheObjMap  redisConn
	time.Sleep(2 * time.Second)
	for _, obj := range CacheObjMap.Keys() {
		//objid:=obj.(string)
		unique, _ := CacheObjMap.Get(obj)
		uniqueid := unique.(string)
		// 没有 objid 存入 所有 数据
		keysMatch, err := RedisCli.Do("keys", fmt.Sprintf("cc:objdata:%v", obj)).Result()
		if err != nil {
			blog.Errorf("sync objid:%s data err:%v", obj, err)
		}

		//blog.Errorf("sync objid keymatch current:%v", objid,keysMatch)
		keyArrs, ok := keysMatch.([]interface{})
		if ok {
			if len(keyArrs) == 0 {
				// mongo中找到所有 iscache的模型数据，跳过在redis 中有 cc:objdata:{objid}开头的
				err, redisSetMap := ObjAllDataFind(obj, uniqueid)
				sliceResult := RedisInter.Keys(context.Background(), fmt.Sprintf(Rediskeyprefix, obj, "*"))
				if len(sliceResult.Val()) > 0 {
					continue
				}
				if err != nil {
					blog.Errorf("sync AllObjData obj:%s  err:%v", obj, err)
					continue
				}
				if len(redisSetMap) == 0 {
					continue
				}
				writeRedis(redisSetMap)

			}
		}
	}
}
func ObjAllDataFind(obj, unique string) (err error, redisSetMap map[string]interface{}) {
	var result []map[string]interface{}
	redisSetMap = make(map[string]interface{}, len(result))
	cond := operation.M{"bk_supplier_account": "0", "bk_obj_id": obj}
	err = mongodb.Client().Table(common.BKTableNameBaseInst).Find(cond).All(context.Background(), &result)
	if err != nil {
		return
	}

	for _, value := range result {
		b, err := json.Marshal(value)
		if err != nil {
			//fmt.Println("json.Marshal failed:", err)
			blog.Errorf("json.Marshal failed:%v", err)
		}
		redisSetMap[fmt.Sprintf(Rediskeyprefix, obj, value[unique])] = b
	}
	return
}

func writeRedis(redisSetMap map[string]interface{}) {
	RedisInter.SetMany(context.Background(), redisSetMap)
	go func() {
		RedisInter.SetManyExpiration(topCtx, redisSetMap, time.Duration(24*time.Hour))
	}()

}
