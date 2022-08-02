package service

import (
	"configcenter/src/common/blog"
	"configcenter/src/common/core/utils"
	"errors"

	"configcenter/src/common/http/rest"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/source_controller/coreservice/bussiness/cache"
	"fmt"
)

func (s *coreService) SearchModelInstancesCache(ctx *rest.Contexts) {
	inputData := make(map[string]interface{}, 0)
	if err := ctx.DecodeInto(&inputData); nil != err {
		ctx.RespAutoError(err)
		return
	}

	objectID := ctx.Request.PathParameter("bk_obj_id")
	//fmt.Println(objectID)

	var iscache bool
	var ui interface{}
	var unique string
	//1. find in iscache obj and unique , 从input 是否存在unique

	ui, iscache = cache.CacheObjMap.Get(objectID)
	// 没有cache 使用mongo
	//if !iscache {
	//	ctx.RespEntityWithError(nil, errors.New(fmt.Sprintf("cahceObjMap not found obj: %s not found  ", objectID)))
	//	return
	//}
	if iscache {
		err, uniqueValue := cache.CacheObjMap.FindKeyAndValue(objectID, inputData)
		if nil != err {
			//ctx.RespEntityWithError()
			blog.Errorf(fmt.Sprintf("obj: %s not found unique err:%v ", objectID, err))
			//return
		}
		unique = ui.(string)
		//2. get redis data
		cacheobj := &cache.CacheData{
			Type:        cache.Select,
			ObjID:       objectID,
			UniqueValue: uniqueValue,
			//Data:          nil,
			//CacheSaveData: nil,
		}
		cacheResult, err := cacheobj.FindOne()
		if nil == err {
			//ctx.RespEntityWithError(fmt.Sprintf("obj: %s unique:%s get cache err:%v ", objectID, unique, err), err)
			resultMap, err := utils.JsonToMap(cacheResult)
			if nil != err {
				ctx.RespEntityWithError(fmt.Sprintf("obj: %s unique:%s JsontoMap err:%v ", objectID, unique, err), err)
				return
			}
			//
			dataResult := &metadata.QueryResult{
				Count: 1,
				Info:  []mapstr.MapStr{resultMap},
			}

			ctx.RespEntityHeader(dataResult, map[string]string{"iscache": "1"})
			return
		}
	}

	//3. not found redis ,find mongo

	//// 4.  save redis
	//// 原则是 只 保存一条到cache
	sc := &cache.SendCache{
		Cond:     inputData,
		Core:     s.core,
		ObjectID: objectID,
	}

	sc.CopyKit(ctx.Kit)
	sc.IsCache()
	// 不判断 cache,mongo 查找 数据
	err, dataResult := sc.SearchDBAndSaveCache()
	if err != nil {
		ctx.RespEntityWithError(fmt.Sprintf("obj: %s unique:%s  get mongo err:%v ", objectID, unique, err), err)
		return
	}

	ctx.RespEntityHeader(dataResult, map[string]string{"iscache": "0"})

}

func (s *coreService) UpdateModelManyInstancesCache(ctx *rest.Contexts) {
	inputData := metadata.UpdateManyModelInstanceMany{}
	if err := ctx.DecodeInto(&inputData); nil != err {
		ctx.RespAutoError(err)
		return
	}
	objectID := ctx.Request.PathParameter("bk_obj_id")
	if objectID == "" || inputData.UniqueKey == "" {
		ctx.RespAutoError(errors.New(fmt.Sprintf("params err")))
		return
	}
	UniqueKey := inputData.UniqueKey

	var successSaveDB = make([]interface{}, 0, len(inputData.Datas))
	var errSaveDB = make([]interface{}, 0, len(inputData.Datas))
	var successSaveDBMap []mapstr.MapStr = make([]mapstr.MapStr, 0, len(inputData.Datas))

	for _, inputRow := range inputData.Datas {
		//fmt.Println(inputRow)
		//inputCopy := inputRow.Condition.Clone()
		if _, ok := inputRow[UniqueKey]; !ok {
			continue
		}
		inputUpdate := metadata.UpdateOption{
			Data:      inputRow,
			Condition: mapstr.MapStr{UniqueKey: inputRow[UniqueKey]},
			//CanEditAll: false,
		}
		dataResult, err := s.core.InstanceOperation().UpdateModelInstance(ctx.Kit, objectID, inputUpdate)
		//fmt.Println(dataResult, err)
		if err == nil && dataResult.Count > 0 {
			successSaveDBMap = append(successSaveDBMap, inputRow)
			successSaveDB = append(successSaveDB, inputRow[UniqueKey])

		} else {
			errSaveDB = append(errSaveDB, inputRow[UniqueKey].(string))
			blog.Errorf("Update Instance Cache save db Cond:[%v] err:%v", inputUpdate, err)
			continue
		}
	}

	var successSaveCache []interface{}
	//var iscache bool
	//var ui interface{}
	//var unique string
	//ui, iscache = cache.CacheObjMap.Get(objectID)
	//unique = ui.(string)
	sc := &cache.SendCache{
		//Cond: inputCopy,
		Core:     s.core,
		ObjectID: objectID,
	}
	sc.IsCache()
	if sc.Iscache {
		successSaveCache = make([]interface{}, 0, len(successSaveDB))
		for _, inputdata := range successSaveDBMap {
			inputCopy := inputdata.Clone()
			sc.Cond = inputCopy

			sc.CopyKit(ctx.Kit)

			err, _ := sc.SearchDBAndSaveCache()
			//fmt.Println(dataResult)
			//fmt.Println(err)
			if err != nil {
				blog.Errorf("Update Instance Cache save Cache Cond:[%v] err:%v", inputdata, err)
				continue
			}
			successSaveCache = append(successSaveCache, inputdata[UniqueKey])
		}
	}

	//fmt.Println(result)
	ctx.RespEntityWithError(map[string]interface{}{"successDB": successSaveDB, "errDB": errSaveDB}, nil)
	//ctx.RespMap(map[string]interface{}{"successDB": successSaveDB, "successCache": successSaveCache})

}
