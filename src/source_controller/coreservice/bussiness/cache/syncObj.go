package cache

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/core/utils"
	"configcenter/src/source_controller/coreservice/core/operation"
	"configcenter/src/storage/driver/mongodb"
	"context"
	"time"
)

/**
 * @Author: Tao Jun
 * @Since: 2022/7/6
 * @Desc: syncunique.go
**/

var CacheObjMap *utils.SafeMap
var cacheMapTicker *time.Ticker
var CacheMapChan chan struct{} = make(chan struct{}, 0)

type syncCacheObj struct {
	objList        []string
	cacheObjUnique map[string]interface{}
}

func init() {
	CacheObjMap = utils.NewSafeMap()
}

func SyncObjCache() {
	cacheMapTicker = time.NewTicker(time.Hour * 1)
	// 初始化 objcachemap
	SyncObjCacheMap()

	// crontab objcachemap
	go func() {
		for {
			select {
			case <-cacheMapTicker.C:
				SyncObjCacheMap()
			}
		}
	}()

	// wait chan<-
	go func() {
		for {
			select {
			case <-CacheMapChan:
				SyncObjCacheMap()
			}
		}
	}()

}

func SyncObjCacheMap() {
	ts := new(syncCacheObj)
	ts.sync()
	ts.writeCachemap()
	blog.Errorf("sync cacheObj plan current:%v", CacheObjMap.M)
}

func (s *syncCacheObj) filterCacheObj(ctx context.Context) (err error) {

	var result []map[string]interface{}
	//var result1 []interface{}
	//err = mongodb.Client().Table(common.BKTableNameObjDes).Find(operation.M{"iscache": true}).
	//Fields("bk_obj_id").All(ctx, &result)
	err = mongodb.Client().Table(common.BKTableNameObjDes).Find(operation.M{}).
		Fields("bk_obj_id").All(ctx, &result)
	//if err != nil {
	//	return err, result
	//}
	var rr []string = make([]string, 0, len(result))
	for _, v := range result {
		rr = append(rr, v["bk_obj_id"].(string))
	}
	//copy(rr, result)
	s.objList = rr
	return err
}

func (s *syncCacheObj) sync() {
	var ctx context.Context = context.Background()
	err := s.filterCacheObj(ctx)
	if err != nil {
		blog.Errorf("sync host cacheObj err:%v", err)
		return
	}
	if len(s.objList) == 0 {
		blog.Errorf("sync host cacheObj len:%v", 0)
		return
	}
	s.matchUniqueDes()

}

func (s *syncCacheObj) matchUniqueDes() {
	s.cacheObjUnique = make(map[string]interface{}, len(s.objList))

	var pipeline []operation.M
	for _, obj := range s.objList {
		pipeline = []operation.M{
			//{"$match": operation.M{"must_check": true, "bk_supplier_account": "0"}}, // 属性为空时检验
			{"$match": operation.M{"must_check": false, "bk_supplier_account": "0"}}, // 属性为空时检验
			{"$unwind": "$keys"},
			{"$project": operation.M{"key_id": "$keys.key_id", "bk_obj_id": "$bk_obj_id"}},
			{"$lookup": operation.M{"from": "cc_ObjAttDes", "localField": "key_id", "foreignField": "id", "as": "AttDes"}},
			{"$unwind": "$AttDes"},
			{"$match": operation.M{"AttDes.bk_property_id": operation.M{"$ne": "bk_inst_name"}}}, // 不使用 默认实例
			{"$project": operation.M{"_id": 0, "key_id": "$key_id", "bk_obj_id": "$bk_obj_id", "bk_property_id": "$AttDes.bk_property_id"}},
			{"$sort": operation.M{"key_id": 1}},
			{"$group": operation.M{"_id": "$bk_obj_id", "first_unique": operation.M{"$first": "$bk_property_id"}}}, //排序后 取第一个 唯一字段
			{"$match": operation.M{"_id": obj}},
		}
		var result map[string]interface{}
		if err := mongodb.Client().Table(common.BKTableNameObjUnique).AggregateOne(context.Background(), pipeline, &result); err != nil {
			//blog.Errorf("sync obj :%v err:%v", obj, err)
			continue
		}
		if len(result) == 0 {
			blog.Errorf("sync obj :%v len is %d", obj, 0)
			continue
		}
		//fmt.Println(result)
		//cacheObjUnique[]
		s.cacheObjUnique[obj] = result["first_unique"]
	}
}

func (s *syncCacheObj) writeCachemap() {
	CacheObjMap.Clear()
	for key, value := range s.cacheObjUnique {
		CacheObjMap.Put(key, value)
	}
}
