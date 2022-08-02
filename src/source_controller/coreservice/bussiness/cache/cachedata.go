package cache

import (
	"configcenter/src/common/blog"
	"configcenter/src/common/http/rest"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/cloud_server/common"
	"configcenter/src/source_controller/coreservice/core"
	"context"
	"errors"
	"fmt"
	"time"
)

/**
 * @Author: Tao Jun
 * @Since: 2022/7/6
 * @Desc: syncunique.go
**/

var topCtx context.Context = context.Background()
var SendCacheChan chan *SendCache = make(chan *SendCache, 2048)

//const redisdataprefix = "cc:objdata:%s"
const Rediskeyprefix = "cc:objdata:%s:%s"
const redisTimeoutSec = 10
const mongoWriteTimeoutSec = 60 * 10

type CacheType string

var (
	Update CacheType = "update"
	//Create CacheType="create"
	Delete CacheType = "delete"
	Select CacheType = "select"
)

// 单条数据
type CacheData struct {
	Type          CacheType
	ObjID         string
	Unique        string
	UniqueValue   string
	Data          []mapstr.MapStr // mongodata
	CacheSaveData map[string]interface{}
	CacheDelKeys  []string
}

type SendCache struct {
	Iscache  bool
	Unique   string
	ObjectID string
	Core     core.Core
	Cond     mapstr.MapStr
	Kit      *rest.Kit
}

func NewCacheData(t CacheType) *CacheData {
	return &CacheData{
		Type: t,
	}
}

func (c *CacheData) SearchOneData() (str string, err error) {

	str, err = RedisCli.Get(fmt.Sprintf(Rediskeyprefix, c.ObjID, c.Unique)).Result()
	return
}

func (c *CacheData) FormatRedisKeyData() {
	result := make(map[string]interface{}, len(c.Data))
	for _, row := range c.Data {
		if uniqueValue, ok := row[c.Unique]; ok {
			if json, err := row.ToJSON(); err == nil {
				result[fmt.Sprintf(Rediskeyprefix, c.ObjID, uniqueValue.(string))] = json
			}

		}
	}
	c.CacheSaveData = result

}

func (c *CacheData) FormatRedisKey() {
	result := make([]string, 0, len(c.Data))
	for _, row := range c.Data {
		if uniqueValue, ok := row[c.Unique]; ok {

			result = append(result, fmt.Sprintf(Rediskeyprefix, c.ObjID, uniqueValue.(string)))

		}
	}
	c.CacheDelKeys = result

}

func (c *CacheData) FindOne() (string, error) {
	if c.Type == Select {
		ctx, _ := context.WithTimeout(topCtx, time.Duration(redisTimeoutSec*time.Second))
		return RedisInter.Get(ctx, fmt.Sprintf(Rediskeyprefix, c.ObjID, c.UniqueValue)).Result()
	} else {
		return "", errors.New(fmt.Sprintf("Type not Select"))
	}

}
func (c *CacheData) SendManyCache() {
	//for _, data := range datas {
	//	SendOneCacheToChan(objid, unique, data.ToMapInterface(), cacheType)
	//}
	CacheDataChan <- c
	//fmt.Println("send many cache")
}

func (s *SendCache) IsCache() {
	var iscache bool
	var ui interface{}
	ui, iscache = CacheObjMap.Get(s.ObjectID)
	if !iscache {
		blog.Infof("objid:%s not found cache skip", s.ObjectID)
		return
	}
	s.Iscache = iscache
	s.Unique = ui.(string)
}

func (s *SendCache) CopyKit(k *rest.Kit) {
	kit := common.NewKit()
	kit.SupplierAccount = k.SupplierAccount
	kit.Header = k.Header
	kit.CCError = k.CCError
	kit.User = k.User
	kit.Rid = k.Rid
	kit.Ctx, _ = context.WithTimeout(topCtx, time.Duration(mongoWriteTimeoutSec*time.Second))
	s.Kit = kit
}

func (s *SendCache) SearchDBAndSaveCache() (error, *metadata.QueryResult) {

	inputMongoData := metadata.QueryCondition{Condition: s.Cond}

	dataResult, err := s.Core.InstanceOperation().SearchModelInstance(s.Kit, s.ObjectID, inputMongoData)
	if nil != err {
		return err, nil
	}
	//fmt.Println(dataResult)
	//fmt.Println(inputMongoData)
	// 4.  save redis
	// 原则是 只 保存一条到cache
	if dataResult.Count > 0 && s.Iscache {
		//SendManyCache(objectID, unique, dataResult.Info, Update)
		cacheData := &CacheData{
			Type:   Update,
			ObjID:  s.ObjectID,
			Unique: s.Unique,
			Data:   dataResult.Info,
			//CacheSaveData: nil,
		}
		cacheData.FormatRedisKeyData()
		//fmt.Printf("%+v\n", cacheData)
		if len(cacheData.CacheSaveData) > 0 {
			cacheData.SendManyCache()
		}
	}
	return nil, dataResult
}

func (s *SendCache) DeleteCache(dataResult *metadata.QueryResult) (error, *metadata.QueryResult) {

	//fmt.Println(dataResult)
	//fmt.Println(inputMongoData)
	// 4.  save redis
	// 原则是 只 保存一条到cache
	if dataResult.Count > 0 && s.Iscache {
		//SendManyCache(objectID, unique, dataResult.Info, Update)
		cacheData := &CacheData{
			Type:   Delete,
			ObjID:  s.ObjectID,
			Unique: s.Unique,
			Data:   dataResult.Info,
			//CacheSaveData: nil,
		}
		cacheData.FormatRedisKey()
		//fmt.Printf("%+v\n", cacheData)
		if len(cacheData.CacheDelKeys) > 0 {
			cacheData.SendManyCache()
		}
	}
	return nil, dataResult
}

func WaitSendCacheData() {

	go func() {
		SendCacheChanAction()
	}()
}

func SendCacheChanAction() {
	for {
		select {
		case sendCacheData := <-SendCacheChan:
			sendCacheData.IsCache()
			if sendCacheData.Iscache {
				err, _ := sendCacheData.SearchDBAndSaveCache()
				if err != nil {
					blog.Errorf("objid:%s unique:%v SendCacheData Chan save redis err:%v", sendCacheData.ObjectID, sendCacheData.Unique, err)
				}
			}

		}
	}
}
