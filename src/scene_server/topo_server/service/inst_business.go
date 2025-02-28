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

package service

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"configcenter/src/ac"
	"configcenter/src/ac/iam"
	"configcenter/src/ac/meta"
	"configcenter/src/common"
	"configcenter/src/common/auth"
	"configcenter/src/common/blog"
	"configcenter/src/common/condition"
	"configcenter/src/common/http/rest"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/mapstruct"
	"configcenter/src/common/metadata"
	gparams "configcenter/src/common/paraparse"
	"configcenter/src/common/querybuilder"
	"configcenter/src/common/util"
	"configcenter/src/scene_server/topo_server/core/inst"
	"configcenter/src/thirdparty/hooks"
)

// CreateBusiness create a new business
func (s *Service) CreateBusiness(ctx *rest.Contexts) {
	data := mapstr.MapStr{}
	if err := ctx.DecodeInto(&data); err != nil {
		ctx.RespAutoError(err)
		return
	}

	if err := hooks.ValidateCreateBusinessHook(ctx.Kit, s.Engine.CoreAPI, data); err != nil {
		blog.Errorf("validate create business hook failed, err: %v, rid: %s", err, ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}

	obj, err := s.Core.ObjectOperation().FindSingleObject(ctx.Kit, common.BKInnerObjIDApp)
	if nil != err {
		blog.Errorf("failed to search the business, %s, rid: %s", err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}

	data.Set(common.BKDefaultField, common.DefaultFlagDefaultValue)
	// do with transaction
	var business inst.Inst
	txnErr := s.Engine.CoreAPI.CoreService().Txn().AutoRunTxn(ctx.Kit.Ctx, ctx.Kit.Header, func() error {
		var err error
		business, err = s.Core.BusinessOperation().CreateBusiness(ctx.Kit, obj, data)
		if err != nil {
			blog.Errorf("create business failed, err: %v, rid: %s", err, ctx.Kit.Rid)
			return err
		}

		// register business resource creator action to iam
		if auth.EnableAuthorize() {
			var bizID int64
			if bizID, err = business.GetBizID(); err != nil {
				blog.ErrorJSON("get biz id failed, err: %s, biz: %s, rid: %s", err, business, ctx.Kit.Rid)
				return err
			}
			var bizName string
			if bizName, err = business.GetInstName(); err != nil {
				blog.ErrorJSON("get biz name failed, err: %s, biz: %s, rid: %s", err, business, ctx.Kit.Rid)
				return err
			}
			iamInstance := metadata.IamInstanceWithCreator{
				Type:    string(iam.Business),
				ID:      strconv.FormatInt(bizID, 10),
				Name:    bizName,
				Creator: ctx.Kit.User,
			}
			_, err = s.AuthManager.Authorizer.RegisterResourceCreatorAction(ctx.Kit.Ctx, ctx.Kit.Header, iamInstance)
			if err != nil {
				blog.Errorf("register created business to iam failed, err: %v, rid: %s", err, ctx.Kit.Rid)
				return err
			}
		}
		return nil
	})

	if txnErr != nil {
		ctx.RespAutoError(txnErr)
		return
	}
	ctx.RespEntity(business)
}

// UpdateBusiness update the business
func (s *Service) UpdateBusiness(ctx *rest.Contexts) {
	data := mapstr.MapStr{}
	if err := ctx.DecodeInto(&data); err != nil {
		ctx.RespAutoError(err)
		return
	}

	obj, err := s.Core.ObjectOperation().FindSingleObject(ctx.Kit, common.BKInnerObjIDApp)
	if nil != err {
		blog.Errorf("failed to search the business, %s, rid: %s", err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}

	bizID, err := strconv.ParseInt(ctx.Request.PathParameter("app_id"), 10, 64)
	if nil != err {
		blog.Errorf("failed to parse the bizId, error: %v, rid: %s", err, ctx.Kit.Rid)
		ctx.RespAutoError(ctx.Kit.CCError.CCErrorf(common.CCErrCommParamsNeedInt, "business id"))
		return
	}

	txnErr := s.Engine.CoreAPI.CoreService().Txn().AutoRunTxn(ctx.Kit.Ctx, ctx.Kit.Header, func() error {
		err = s.Core.BusinessOperation().UpdateBusiness(ctx.Kit, data, obj, bizID)
		if err != nil {
			return err
		}
		return nil
	})

	if txnErr != nil {
		ctx.RespAutoError(txnErr)
		return
	}
	ctx.RespEntity(nil)
}

// UpdateBusinessStatus update the business status
func (s *Service) UpdateBusinessStatus(ctx *rest.Contexts) {
	data := struct {
		metadata.UpdateBusinessStatusOption `json:",inline"`
	}{}
	if err := ctx.DecodeInto(&data); err != nil {
		ctx.RespAutoError(err)
		return
	}

	obj, err := s.Core.ObjectOperation().FindSingleObject(ctx.Kit, common.BKInnerObjIDApp)
	if nil != err {
		blog.Errorf("failed to search the business, %s, rid: %s", err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}

	bizID, err := strconv.ParseInt(ctx.Request.PathParameter("app_id"), 10, 64)
	if nil != err {
		blog.Errorf("[api-business]failed to parse the biz id, error: %s, rid: %s", err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(ctx.Kit.CCError.CCErrorf(common.CCErrCommParamsNeedInt, "business id"))
		return
	}
	query := &metadata.QueryCondition{
		Condition: mapstr.MapStr{common.BKAppIDField: bizID},
	}
	_, bizs, err := s.Core.BusinessOperation().FindBiz(ctx.Kit, query)
	if len(bizs) <= 0 {
		ctx.RespAutoError(ctx.Kit.CCError.CCError(common.CCErrCommNotFound))
		return
	}
	biz := metadata.BizBasicInfo{}
	if err := mapstruct.Decode2Struct(bizs[0], &biz); err != nil {
		blog.Errorf("[api-business]failed, parse biz failed, biz: %+v, err: %s, rid: %s", bizs[0], err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(ctx.Kit.CCError.CCError(common.CCErrCommParseDBFailed))
		return
	}

	updateData := mapstr.New()
	switch common.DataStatusFlag(ctx.Request.PathParameter("flag")) {
	case common.DataStatusDisabled:
		if err := s.Core.AssociationOperation().CheckAssociation(ctx.Kit, obj.Object().ObjectID, bizID); nil != err {
			ctx.RespAutoError(err)
			return
		}

		// check if this business still has hosts.
		has, err := s.Core.BusinessOperation().HasHosts(ctx.Kit, bizID)
		if err != nil {
			ctx.RespAutoError(err)
			return
		}
		if has {
			ctx.RespAutoError(ctx.Kit.CCError.CCError(common.CCErrTopoArchiveBusinessHasHost))
			return
		}
		achieveBizName, err := s.Core.BusinessOperation().GenerateAchieveBusinessName(ctx.Kit, biz.BizName)
		if err != nil {
			ctx.RespAutoError(err)
			return
		}
		updateData.Set(common.BKAppNameField, achieveBizName)
		updateData.Set(common.BKDataStatusField, ctx.Request.PathParameter("flag"))
	case common.DataStatusEnable:
		if len(data.UpdateBusinessStatusOption.BizName) > 0 {
			updateData.Set(common.BKAppNameField, data.UpdateBusinessStatusOption.BizName)
		}
		updateData.Set(common.BKDataStatusField, ctx.Request.PathParameter("flag"))
	default:
		ctx.RespAutoError(ctx.Kit.CCError.CCErrorf(common.CCErrCommParamsIsInvalid, ctx.Request.PathParameter))
		return
	}

	txnErr := s.Engine.CoreAPI.CoreService().Txn().AutoRunTxn(ctx.Kit.Ctx, ctx.Kit.Header, func() error {
		err = s.Core.BusinessOperation().UpdateBusiness(ctx.Kit, updateData, obj, bizID)
		if err != nil {
			blog.Errorf("UpdateBusinessStatus failed, run update failed, err: %+v, rid: %s", err, ctx.Kit.Rid)
			return err
		}
		return nil
	})

	if txnErr != nil {
		ctx.RespAutoError(txnErr)
		return
	}
	ctx.RespEntity(nil)
}

// UpdateBizPropertyBatch batch update business properties
func (s *Service) UpdateBizPropertyBatch(ctx *rest.Contexts) {
	param := new(metadata.UpdateBizPropertyBatchParameter)
	if err := ctx.DecodeInto(param); err != nil {
		ctx.RespAutoError(err)
		return
	}

	obj, err := s.Core.ObjectOperation().FindSingleObject(ctx.Kit, common.BKInnerObjIDApp)
	if nil != err {
		blog.Errorf("failed to search the business, %s, rid: %s", err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}

	bizIDs, err := s.getBizIDByCond(ctx, param.Condition)
	if err != nil {
		ctx.RespAutoError(err)
		return
	}

	if len(bizIDs) <= 0 {
		blog.Errorf("found no business by condition, rid: %s", ctx.Kit.Rid)
		ctx.RespAutoError(ctx.Kit.CCError.CCError(common.CCErrCommNotFound))
		return
	}

	// check authorization
	if s.AuthManager.Enabled() {
		if err := s.AuthManager.AuthorizeByBusinessID(ctx.Kit.Ctx, ctx.Kit.Header, meta.Update, bizIDs...); err ==
			ac.NoAuthorizeError {
			perm, err := s.AuthManager.GenBizBatchNoPermissionResp(ctx.Kit.Ctx, ctx.Kit.Header, meta.Update, bizIDs)
			if err != nil && err != ac.NoAuthorizeError {
				blog.Errorf("get biz permission failed, biz: %v, err: %v, rid: %s", bizIDs, err, ctx.Kit.Rid)
				ctx.RespAutoError(ctx.Kit.CCError.CCError(common.CCErrCommAuthorizeFailed))
				return
			}
			ctx.RespEntityWithError(perm, ac.NoAuthorizeError)
			return
		} else if err != nil {
			blog.Errorf("biz authorize failed, biz: %v, err: %v, rid: %s", bizIDs, err, ctx.Kit.Rid)
			ctx.RespAutoError(ctx.Kit.CCError.CCError(common.CCErrCommAuthorizeFailed))
			return
		}
	}

	updateCond := mapstr.MapStr{
		common.BKAppIDField: mapstr.MapStr{
			common.BKDBIN: bizIDs,
		},
		// exclude archived biz
		common.BKDataStatusField: mapstr.MapStr{
			common.BKDBNE: common.DataStatusDisabled,
		},
		// exclude default biz
		common.BKDefaultField: mapstr.MapStr{
			common.BKDBNE: common.DefaultAppFlag,
		},
	}

	data := param.Properties
	// cannot update biz bk_data_status
	delete(data, common.BKDataStatusField)

	// update biz instances
	txnErr := s.Engine.CoreAPI.CoreService().Txn().AutoRunTxn(ctx.Kit.Ctx, ctx.Kit.Header, func() error {
		if err := s.Core.BusinessOperation().UpdateBusinessByCond(ctx.Kit, data, obj, updateCond); err != nil {
			blog.Errorf("update many biz err: %+v, rid: %s", err, ctx.Kit.Rid)
			return err
		}
		return nil
	})

	if txnErr != nil {
		ctx.RespAutoError(txnErr)
		return
	}
	ctx.RespEntity(nil)
}

func (s *Service) getBizIDByCond(ctx *rest.Contexts, cond mapstr.MapStr) ([]int64, error) {
	attrCond := condition.CreateCondition()
	attrCond.Field(metadata.AttributeFieldObjectID).Eq(common.BKInnerObjIDApp)
	attrCond.Field(metadata.AttributeFieldPropertyType).Eq(common.FieldTypeUser)
	attrArr, err := s.Core.AttributeOperation().FindBusinessAttribute(ctx.Kit, attrCond.ToMapStr())
	if nil != err {
		blog.Errorf("failed get the business attribute, %s, rid:%s", err.Error(), ctx.Kit.Rid)
		return nil, err
	}
	// userFieldArr Fields in the business are user-type fields
	var userFields []string
	for _, attribute := range attrArr {
		userFields = append(userFields, attribute.PropertyID)
	}

	filter := handleSpecialBusinessFieldSearchCond(cond, userFields)
	// can only find normal business, but not resource pool business
	filter[common.BKDefaultField] = mapstr.MapStr{
		common.BKDBNE: common.DefaultAppFlag,
	}

	distinctOpt := &metadata.DistinctFieldOption{
		TableName: common.BKTableNameBaseApp,
		Field:     common.BKAppIDField,
		Filter:    filter,
	}

	rst, err := s.Engine.CoreAPI.CoreService().Common().GetDistinctField(ctx.Kit.Ctx, ctx.Kit.Header, distinctOpt)
	if err != nil {
		blog.Errorf("get biz ids failed, distinct opt: %+v, err: %v, rid: %s", distinctOpt, err, ctx.Kit.Rid)
		return nil, err
	}

	bizIDs, err := util.SliceInterfaceToInt64(rst)
	if err != nil {
		blog.Errorf("biz ids to int failed, biz ids: %v, err: %v, rid: %s", rst, err, ctx.Kit.Rid)
		return nil, err
	}

	return bizIDs, nil
}

// DeleteBusiness delete archived business
func (s *Service) DeleteBusiness(ctx *rest.Contexts) {
	param := new(metadata.DeleteBizParam)
	if err := ctx.DecodeInto(param); err != nil {
		ctx.RespAutoError(err)
		return
	}

	if len(param.BizID) == 0 {
		blog.Errorf("invalid bk_biz_id len 0, rid: %s", ctx.Kit.Rid)
		ctx.RespAutoError(ctx.Kit.CCError.CCErrorf(common.CCErrCommParamsIsInvalid, common.BKAppIDField))
		return
	}

	if len(param.BizID) > common.BKDefaultLimit {
		blog.Errorf("bk_biz_id len %d exceed max page size %d, rid:%s", len(param.BizID),
			common.BKDefaultLimit, ctx.Kit.Rid)
		ctx.RespAutoError(ctx.Kit.CCError.CCErrorf(common.CCErrCommXXExceedLimit, "update",
			common.BKDefaultLimit))
		return
	}

	if err := hooks.ValidateDeleteBusinessHook(ctx.Kit, s.Engine.CoreAPI, param.BizID); err != nil {
		blog.Errorf("validate delete business hook failed, err: %v, rid: %s", err, ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}

	// delete biz instances and related resources
	txnErr := s.Engine.CoreAPI.CoreService().Txn().AutoRunTxn(ctx.Kit.Ctx, ctx.Kit.Header, func() error {
		if err := s.Core.BusinessOperation().DeleteBusiness(ctx.Kit, param.BizID); err != nil {
			blog.Errorf("failed to delete biz, ids: %v, err: %v, rid: %s", param.BizID, err, ctx.Kit.Rid)
			return err
		}
		return nil
	})

	if txnErr != nil {
		ctx.RespAutoError(txnErr)
		return
	}
	ctx.RespEntity(nil)
}

// find business list with these info：
// 1. have any authorized resources in a business.
// 2. only returned with a few field for this business info.
func (s *Service) SearchReducedBusinessList(ctx *rest.Contexts) {
	page := metadata.BasePage{
		Limit: common.BKNoLimit,
	}
	sortParam := ctx.Request.QueryParameter("sort")
	if len(sortParam) > 0 {
		page.Sort = sortParam
	}
	if errKey, err := page.Validate(true); err != nil {
		blog.Errorf("page parameter invalid, errKey: %v, err: %s, rid: %s", errKey, err, ctx.Kit.Rid)
		ctx.RespAutoError(ctx.Kit.CCError.CCErrorf(common.CCErrCommParamsInvalid, errKey))
		return
	}
	query := &metadata.QueryCondition{
		Fields: []string{common.BKAppIDField, common.BKAppNameField},
		Page:   page,
		Condition: mapstr.MapStr{
			common.BKDataStatusField: mapstr.MapStr{common.BKDBNE: common.DataStatusDisabled},
			common.BKDefaultField:    0,
		},
	}

	if s.AuthManager.Enabled() {
		authInput := meta.ListAuthorizedResourcesParam{
			UserName:     ctx.Kit.User,
			ResourceType: meta.Business,
			Action:       meta.ViewBusinessResource,
		}
		authorizedRes, err := s.AuthManager.Authorizer.ListAuthorizedResources(ctx.Kit.Ctx, ctx.Kit.Header, authInput)
		if err != nil {
			blog.Errorf("[api-biz] SearchReducedBusinessList failed, ListAuthorizedResources failed, user: %s,"+
				" err: %s, rid: %s", ctx.Kit.User, err.Error(), ctx.Kit.Rid)
			ctx.RespAutoError(ctx.Kit.CCError.CCError(common.CCErrorTopoGetAuthorizedBusinessListFailed))
			return
		}
		// if isAny is false,we should add bizIds condition.
		if !authorizedRes.IsAny {
			appList := make([]int64, 0)
			for _, resourceID := range authorizedRes.Ids {
				bizID, err := strconv.ParseInt(resourceID, 10, 64)
				if err != nil {
					blog.Errorf("parse bizID(%s) failed, err: %v, rid: %s", bizID, err, ctx.Kit.Rid)
					ctx.RespAutoError(ctx.Kit.CCError.CCErrorf(common.CCErrCommParamsNeedInt, common.BKAppIDField))
					return
				}
				appList = append(appList, bizID)
			}
			if len(appList) == 0 {
				ctx.RespEntityWithCount(0, make([]mapstr.MapStr, 0))
				return
			}
			// sort for prepare to find business with page.
			sort.Sort(util.Int64Slice(appList))
			// user can only find business that is already authorized.
			query.Condition[common.BKAppIDField] = mapstr.MapStr{common.BKDBIN: appList}
		}
	}

	cnt, instItems, err := s.Core.BusinessOperation().FindBiz(ctx.Kit, query)
	if nil != err {
		blog.Errorf("[api-business] failed to find the objects(%s), error info is %s, rid: %s",
			ctx.Request.PathParameter("obj_id"), err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}

	datas := make([]mapstr.MapStr, 0)
	for _, item := range instItems {
		inst := mapstr.New()
		inst[common.BKAppIDField] = item[common.BKAppIDField]
		inst[common.BKAppNameField] = item[common.BKAppNameField]
		datas = append(datas, inst)
	}

	result := mapstr.MapStr{}
	result.Set("count", cnt)
	result.Set("info", datas)
	ctx.RespEntity(result)
}

func (s *Service) GetBusinessBasicInfo(ctx *rest.Contexts) {
	bizID, err := strconv.ParseInt(ctx.Request.PathParameter("app_id"), 10, 64)
	if nil != err {
		blog.Errorf("[api-business]failed to parse the biz id, error: %s, rid: %s", err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(ctx.Kit.CCError.CCErrorf(common.CCErrCommParamsNeedInt, "business id"))
		return
	}
	query := &metadata.QueryCondition{
		Fields: []string{common.BKAppNameField, common.BKAppIDField},
		Condition: map[string]interface{}{
			common.BKAppIDField: bizID,
		},
	}
	result, err := s.Engine.CoreAPI.CoreService().Instance().ReadInstance(ctx.Kit.Ctx, ctx.Kit.Header,
		common.BKInnerObjIDApp, query)
	if err != nil {
		blog.Errorf("failed to get business by id, bizID: %s, err: %s, rid: %s", bizID, err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}
	if len(result.Data.Info) == 0 {
		blog.Errorf(" get business by id not found, bizID: %d, rid: %s", bizID, ctx.Kit.Rid)
		err := ctx.Kit.CCError.CCError(common.CCErrCommNotFound)
		ctx.RespAutoError(err)
		return
	}
	bizData := result.Data.Info[0]
	ctx.RespEntity(bizData)
}

// 4 scenarios, such as user's name user1, scenarios as follows:
// user1
// user1,user3
// user2,user1
// user2,user1,user4
const exactUserRegexp = `(^USER_PLACEHOLDER$)|(^USER_PLACEHOLDER[,]{1})|([,]{1}USER_PLACEHOLDER[,]{1})|([,]{1}USER_PLACEHOLDER$)`

func handleSpecialBusinessFieldSearchCond(input map[string]interface{}, userFieldArr []string) map[string]interface{} {
	output := make(map[string]interface{})
	exactAnd := make([]map[string]interface{}, 0)
	for i, j := range input {
		if j == nil {
			output[i] = j
			continue
		}

		objType := reflect.TypeOf(j)
		switch objType.Kind() {
		case reflect.String:
			if _, ok := j.(json.Number); ok {
				output[i] = j
				continue
			}
			targetStr := j.(string)
			if util.InStrArr(userFieldArr, i) {
				for _, user := range strings.Split(strings.Trim(targetStr, ","), ",") {
					// search with exactly the user's name with regexpF
					like := strings.Replace(exactUserRegexp, "USER_PLACEHOLDER", gparams.SpecialCharChange(user), -1)
					exactAnd = append(exactAnd, mapstr.MapStr{i: mapstr.MapStr{common.BKDBLIKE: like}})
				}
			} else {
				attrVal := gparams.SpecialCharChange(targetStr)
				output[i] = map[string]interface{}{common.BKDBLIKE: attrVal, common.BKDBOPTIONS: "i"}
			}
		default:
			output[i] = j
		}
	}

	if len(exactAnd) > 0 {
		output[common.BKDBAND] = exactAnd
	}

	return output
}

// SearchBusiness search the business by condition
func (s *Service) SearchBusiness(ctx *rest.Contexts) {
	searchCond := new(metadata.QueryBusinessRequest)
	if err := ctx.DecodeInto(&searchCond); nil != err {
		blog.Errorf("failed to parse the params, error info is %s, rid: %s", err.Error(), ctx.Kit.Rid)
		ctx.RespErrorCodeOnly(common.CCErrCommJSONUnmarshalFailed, "")
		return
	}

	// parameters condition and biz_property_filter cannot be set at the same time.
	if searchCond.Condition != nil && searchCond.BizPropertyFilter != nil {
		blog.Errorf("failed to parse the params, condition and biz_property_filter cannot be set at the same "+
			"time, rid: %s", ctx.Kit.Rid)
		ctx.RespErrorCodeOnly(common.CCErrCommParamsInvalid, "condition and biz_property_filter cannot be set "+
			"at the same time")
		return
	}

	attrCond := condition.CreateCondition()
	attrCond.Field(metadata.AttributeFieldObjectID).Eq(common.BKInnerObjIDApp)
	attrCond.Field(metadata.AttributeFieldPropertyType).Eq(common.FieldTypeUser)
	attrArr, err := s.Core.AttributeOperation().FindBusinessAttribute(ctx.Kit, attrCond.ToMapStr())
	if nil != err {
		blog.Errorf("failed get the business attribute, %s, rid:%s", err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}
	// userFieldArr Fields in the business are user-type fields
	var userFields []string
	for _, attribute := range attrArr {
		userFields = append(userFields, attribute.PropertyID)
	}

	searchCond.Condition = handleSpecialBusinessFieldSearchCond(searchCond.Condition, userFields)

	// parse business id from user's condition for testing.
	var bizIDs, authBizIDs []int64
	defErr := ctx.Kit.CCError
	biz, exist := searchCond.Condition[common.BKAppIDField]
	if exist {
		// constrict that bk_biz_id field can only be a numeric value,
		// operators like or/in/and is not allowed.
		if bizcond, ok := biz.(map[string]interface{}); ok {
			if cond, ok := bizcond["$eq"]; ok {
				bizID, err := util.GetInt64ByInterface(cond)
				if err != nil {
					ctx.RespErrorCodeOnly(common.CCErrCommParamsInvalid, "", common.BKAppIDField)
					return
				}
				bizIDs = []int64{bizID}
			}
			if cond, ok := bizcond["$in"]; ok {
				if conds, ok := cond.([]interface{}); ok {
					for _, c := range conds {
						bizID, err := util.GetInt64ByInterface(c)
						if err != nil {
							ctx.RespErrorCodeOnly(common.CCErrCommParamsInvalid, "", common.BKAppIDField)
							return
						}
						bizIDs = append(bizIDs, bizID)
					}
				}
			}
		} else {
			bizID, err := util.GetInt64ByInterface(searchCond.Condition[common.BKAppIDField])
			if err != nil {
				ctx.RespAutoError(ctx.Kit.CCError.Errorf(common.CCErrCommParamsInvalid, common.BKAppIDField))
				return
			}
			bizIDs = []int64{bizID}
		}
	}

	if s.AuthManager.Enabled() {
		authInput := meta.ListAuthorizedResourcesParam{
			UserName:     ctx.Kit.User,
			ResourceType: meta.Business,
			Action:       meta.Find,
		}
		authorizedRes, err := s.AuthManager.Authorizer.ListAuthorizedResources(ctx.Kit.Ctx, ctx.Kit.Header, authInput)
		if err != nil {
			blog.Errorf("[api-biz] SearchBusiness failed, ListAuthorizedResources failed, user: %s, err: %s,"+
				" rid: %s", ctx.Kit.User, err.Error(), ctx.Kit.Rid)
			ctx.RespErrorCodeOnly(common.CCErrorTopoGetAuthorizedBusinessListFailed, "")
			return
		}
		appList := make([]int64, 0)

		// if isAny is true means we have all bizIds authority, else we should parse ids list that we have authority.
		if authorizedRes.IsAny {
			// if user assign the ids,add the ids to the condition.
			if len(bizIDs) > 0 {
				searchCond.Condition[common.BKAppIDField] = mapstr.MapStr{common.BKDBIN: bizIDs}
			}

		} else {
			for _, resourceID := range authorizedRes.Ids {
				bizID, err := strconv.ParseInt(resourceID, 10, 64)
				if err != nil {
					blog.Errorf("parse bizID(%s) failed, err: %v, rid: %s", bizID, err, ctx.Kit.Rid)
					ctx.RespAutoError(ctx.Kit.CCError.CCErrorf(common.CCErrCommParamsNeedInt, common.BKAppIDField))
					return
				}
				appList = append(appList, bizID)
			}
			if len(bizIDs) > 0 {
				// this means that user want to find a specific business.now we check if he has this authority.
				for _, bizID := range bizIDs {
					if util.InArray(bizID, appList) {
						// authBizIDs store the authorized bizIDs
						authBizIDs = append(authBizIDs, bizID)
					}
				}
				if len(authBizIDs) > 0 {
					searchCond.Condition[common.BKAppIDField] = mapstr.MapStr{common.BKDBIN: authBizIDs}
				} else {
					// if there are no qualified bizIDs, return null
					result := mapstr.MapStr{}
					result.Set("count", 0)
					result.Set("info", []mapstr.MapStr{})
					ctx.RespEntity(result)
					return
				}
				// now you have the authority.
			} else {
				if len(appList) == 0 {
					ctx.RespEntityWithCount(0, make([]mapstr.MapStr, 0))
					return
				}

				// sort for prepare to find business with page.
				sort.Sort(util.Int64Slice(appList))
				// user can only find business that is already authorized.
				searchCond.Condition[common.BKAppIDField] = mapstr.MapStr{common.BKDBIN: appList}
			}
		}
	}

	// Only one of biz_property_filter and condition parameters can take effect, and condition is not recommended to
	// continue to use it.
	if searchCond.BizPropertyFilter != nil {
		if key, err := searchCond.BizPropertyFilter.Validate(); err != nil {
			blog.Errorf("bizPropertyFilter is illegal, err: %v, rid:%s", err, ctx.Kit.Rid)
			ccErr := defErr.CCErrorf(common.CCErrCommParamsInvalid, fmt.Sprintf("biz.property.%s", key))
			ctx.RespAutoError(ccErr)
			return
		}
		if searchCond.BizPropertyFilter.GetDeep() > querybuilder.MaxDeep {
			blog.Errorf("bizPropertyFilter is illegal, err: %v, rid: %s", err, ctx.Kit.Rid, ctx.Kit.Rid)
			ccErr := defErr.CCErrorf(common.CCErrCommParamsInvalid,
				fmt.Sprintf("exceed max query condition deepth: %d", querybuilder.MaxDeep))
			ctx.RespAutoError(ccErr)
			return
		}

		mgoFilter, key, err := searchCond.BizPropertyFilter.ToMgo()
		if err != nil {
			blog.Errorf("BizPropertyFilter ToMgo failed: %s, err: %v,  rid:%s", searchCond.BizPropertyFilter,
				err, ctx.Kit.Rid)
			ctx.RespAutoError(ctx.Kit.CCError.CCErrorf(common.CCErrCommParamsInvalid,
				err.Error()+fmt.Sprintf(", biz_property_filter.%s", key)))
			return
		}
		searchCond.Condition = mgoFilter
	}

	if _, ok := searchCond.Condition[common.BKDataStatusField]; !ok {
		searchCond.Condition[common.BKDataStatusField] = mapstr.MapStr{common.BKDBNE: common.DataStatusDisabled}
	}

	// can only find normal business, but not resource pool business
	searchCond.Condition[common.BKDefaultField] = 0

	query := &metadata.QueryCondition{
		Condition: searchCond.Condition,
		Fields:    searchCond.Fields,
		Page:      searchCond.Page,
	}

	cnt, instItems, err := s.Core.BusinessOperation().FindBiz(ctx.Kit, query)
	if nil != err {
		blog.Errorf("find business failed, err: %v, rid: %s", err, ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}

	result := mapstr.MapStr{}
	result.Set("count", cnt)
	result.Set("info", instItems)

	ctx.RespEntity(result)
}

// SearchOwnerResourcePoolBusiness search archived business by condition
func (s *Service) SearchOwnerResourcePoolBusiness(ctx *rest.Contexts) {

	supplierAccount := ctx.Request.PathParameter("owner_id")
	query := metadata.QueryCondition{
		Condition: mapstr.MapStr{
			common.BKDefaultField:    common.DefaultAppFlag,
			common.BkSupplierAccount: supplierAccount,
		},
	}

	cnt, instItems, err := s.Core.BusinessOperation().FindBiz(ctx.Kit, &query)
	if nil != err {
		blog.Errorf("[api-business] failed to find the objects(%s), error info is %s, rid: %s",
			ctx.Request.PathParameter("obj_id"), err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}
	if cnt == 0 {
		blog.InfoJSON("cond:%s, header:%s, rid:%s", query, ctx.Kit.Header, ctx.Kit.Rid)
	}
	result := mapstr.MapStr{}
	result.Set("count", cnt)
	result.Set("info", instItems)
	ctx.RespEntity(result)
	return
}

// CreateDefaultBusiness create the default business
func (s *Service) CreateDefaultBusiness(ctx *rest.Contexts) {
	data := mapstr.MapStr{}
	if err := ctx.DecodeInto(&data); err != nil {
		ctx.RespAutoError(err)
		return
	}
	obj, err := s.Core.ObjectOperation().FindSingleObject(ctx.Kit, common.BKInnerObjIDApp)
	if nil != err {
		blog.Errorf("failed to search the business, %s, rid: %s", err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}

	data.Set(common.BKDefaultField, common.DefaultAppFlag)

	var business inst.Inst
	txnErr := s.Engine.CoreAPI.CoreService().Txn().AutoRunTxn(ctx.Kit.Ctx, ctx.Kit.Header, func() error {
		var err error
		business, err = s.Core.BusinessOperation().CreateBusiness(ctx.Kit, obj, data)
		if err != nil {
			blog.Errorf("create business failed, err: %+v", err)
			return err
		}
		return nil
	})

	if txnErr != nil {
		ctx.RespAutoError(txnErr)
		return
	}
	ctx.RespEntity(business)
}

func (s *Service) GetInternalModule(ctx *rest.Contexts) {
	bizID, err := strconv.ParseInt(ctx.Request.PathParameter("app_id"), 10, 64)
	if nil != err {
		ctx.RespAutoError(ctx.Kit.CCError.New(common.CCErrTopoAppSearchFailed, err.Error()))
		return
	}

	ctx.SetReadPreference(common.SecondaryPreferredMode)

	_, result, err := s.Core.BusinessOperation().GetInternalModule(ctx.Kit, bizID)
	if nil != err {
		ctx.RespAutoError(err)
		return
	}

	ctx.RespEntity(result)
}

func (s *Service) GetInternalModuleWithStatistics(ctx *rest.Contexts) {
	bizID, err := strconv.ParseInt(ctx.Request.PathParameter("app_id"), 10, 64)
	if nil != err {
		ctx.RespAutoError(ctx.Kit.CCError.New(common.CCErrTopoAppSearchFailed, err.Error()))
		return
	}

	_, innerAppTopo, err := s.Core.BusinessOperation().GetInternalModule(ctx.Kit, bizID)
	if nil != err {
		ctx.RespAutoError(err)
		return
	}
	if innerAppTopo == nil {
		blog.ErrorJSON("GetInternalModuleWithStatistics failed, GetInternalModule return unexpected type: %s,"+
			" rid: %s", innerAppTopo, ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}
	moduleIDArr := make([]int64, 0)
	for _, item := range innerAppTopo.Module {
		moduleIDArr = append(moduleIDArr, item.ModuleID)
	}

	// count host apply rules
	listApplyRuleOption := metadata.ListHostApplyRuleOption{
		ModuleIDs: moduleIDArr,
		Page: metadata.BasePage{
			Limit: common.BKNoLimit,
		},
	}
	hostApplyRules, err := s.Engine.CoreAPI.CoreService().HostApplyRule().ListHostApplyRule(ctx.Kit.Ctx, ctx.Kit.Header,
		bizID, listApplyRuleOption)
	if err != nil {
		blog.Errorf("fillStatistics failed, ListHostApplyRule failed, bizID: %d, option: %+v, err: %+v, rid: %s",
			bizID, listApplyRuleOption, err, ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}
	moduleRuleCount := make(map[int64]int64)
	for _, item := range hostApplyRules.Info {
		if _, exist := moduleRuleCount[item.ModuleID]; exist == false {
			moduleRuleCount[item.ModuleID] = 0
		}
		moduleRuleCount[item.ModuleID] += 1
	}

	set := mapstr.NewFromStruct(innerAppTopo, "field")
	modules := make([]mapstr.MapStr, 0)
	for _, module := range innerAppTopo.Module {
		moduleItem := mapstr.NewFromStruct(module, "field")
		moduleItem["host_apply_rule_count"] = 0
		if ruleCount, ok := moduleRuleCount[module.ModuleID]; ok == true {
			moduleItem["host_apply_rule_count"] = ruleCount
		}
		modules = append(modules, moduleItem)
	}
	set["module"] = modules
	ctx.RespEntity(set)
}

// ListAllBusinessSimplify list all businesses with return only several fields
func (s *Service) ListAllBusinessSimplify(ctx *rest.Contexts) {
	page := metadata.BasePage{
		Limit: common.BKNoLimit,
	}
	sortParam := ctx.Request.QueryParameter("sort")
	if len(sortParam) > 0 {
		page.Sort = sortParam
	}
	if errKey, err := page.Validate(true); err != nil {
		blog.Errorf("[api-biz] ListAllBusinessSimplify failed, page parameter invalid, errKey: %s, err: %s, "+
			"rid: %s", errKey, err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(ctx.Kit.CCError.CCErrorf(common.CCErrCommParamsInvalid, errKey))
		return
	}

	fields := []string{
		common.BKAppIDField,
		common.BKAppNameField,
	}

	query := &metadata.QueryCondition{
		Fields: fields,
		Page:   page,
		Condition: mapstr.MapStr{
			common.BKDataStatusField: mapstr.MapStr{common.BKDBNE: common.DataStatusDisabled},
		},
	}
	cnt, instItems, err := s.Core.BusinessOperation().FindBiz(ctx.Kit, query)
	if nil != err {
		blog.Errorf("ListAllBusinessSimplify failed, FindBusiness failed, err: %s, rid: %s", err.Error(), ctx.Kit.Rid)
		ctx.RespAutoError(err)
		return
	}
	businesses := make([]metadata.BizBasicInfo, 0)
	for _, item := range instItems {
		business := metadata.BizBasicInfo{}
		if err := mapstruct.Decode2Struct(item, &business); err != nil {
			blog.Errorf("ListAllBusinessSimplify failed, decode biz from db failed, err: %s, rid: %s", err.Error(), ctx.Kit.Rid)
			ctx.RespAutoError(ctx.Kit.CCError.CCError(common.CCErrCommParseDBFailed))
			return
		}
		businesses = append(businesses, business)
	}

	result := map[string]interface{}{
		"count": cnt,
		"info":  businesses,
	}
	ctx.RespEntity(result)
}
