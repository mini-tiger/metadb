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
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/common/querybuilder"
	"configcenter/src/common/util"
	webCommon "configcenter/src/web_server/common"
	"configcenter/src/web_server/logics"

	"github.com/gin-gonic/gin"
	"github.com/rentiansheng/xlsx"
)

// ImportHost import host
func (s *Service) ImportHost(c *gin.Context) {
	rid := util.GetHTTPCCRequestID(c.Request.Header)
	ctx := util.NewContextFromHTTPHeader(c.Request.Header)

	language := webCommon.GetLanguageByHTTPRequest(c)
	defLang := s.Language.CreateDefaultCCLanguageIf(language)
	defErr := s.CCErr.CreateDefaultCCErrorIf(language)
	file, err := c.FormFile("file")
	if nil != err {
		blog.Errorf("ImportHost failed, get file from form data failed, err: %+v, rid: %s", err, rid)
		msg := getReturnStr(common.CCErrWebFileNoFound, defErr.Error(common.CCErrWebFileNoFound).Error(), nil)
		c.String(http.StatusOK, msg)
		return
	}

	moduleID := int64(0)
	if moduleIDStr := c.PostForm(common.BKModuleIDField); moduleIDStr != "" {
		moduleID, err = strconv.ParseInt(moduleIDStr, 10, 64)
		if err != nil {
			blog.Errorf("ImportHost failed, bk_module_id not integer, err: %+v, bk_module_id: %s,  rid: %s", err, moduleIDStr, rid)
			msg := getReturnStr(common.CCErrCommParamsNeedInt, defErr.CCErrorf(common.CCErrCommParamsNeedInt, common.BKModuleIDField).Error(), nil)
			c.String(http.StatusOK, msg)
			return
		}
	}

	webCommon.SetProxyHeader(c)

	randNum := rand.Uint32()
	dir := webCommon.ResourcePath + "/import/"
	_, err = os.Stat(dir)
	if nil != err {
		if err := os.MkdirAll(dir, os.ModeDir|os.ModePerm); err != nil {
			blog.Errorf("ImportHost failed, save form data to local file failed, mkdir failed, err: %+v, rid: %s", err, rid)
			c.String(http.StatusInternalServerError, fmt.Sprintf("save form data to local file failed, mkdir failed, err: %+v", err))
			return
		}
	}
	filePath := fmt.Sprintf("%s/importhost-%d-%d.xlsx", dir, time.Now().UnixNano(), randNum)
	if err := c.SaveUploadedFile(file, filePath); nil != err {
		blog.Errorf("ImportHost failed, save form data to local file failed, save data as excel failed, err: %+v, rid: %s", err, rid)
		msg := getReturnStr(common.CCErrWebFileSaveFail, defErr.Errorf(common.CCErrWebFileSaveFail, err.Error()).Error(), nil)
		c.String(http.StatusOK, msg)
		return
	}

	// del file
	defer func(filePath string, rid string) {
		if err := os.Remove(filePath); err != nil {
			blog.Errorf("ImportHost, remove temporary file failed, err: %+v, rid: %s", err, rid)
		}
	}(filePath, rid)

	f, err := xlsx.OpenFile(filePath)
	if nil != err {
		blog.Errorf("ImportHost failed, open form data as excel file failed, err: %+v, rid: %s", err, rid)
		msg := getReturnStr(common.CCErrWebOpenFileFail, defErr.Errorf(common.CCErrWebOpenFileFail, err.Error()).Error(), nil)
		c.String(http.StatusOK, msg)
		return
	}
	result := s.Logics.ImportHosts(ctx, f, c.Request.Header, defLang, 0, moduleID)

	c.JSON(http.StatusOK, result)
}

// ExportHost export host
func (s *Service) ExportHost(c *gin.Context) {
	rid := util.GetHTTPCCRequestID(c.Request.Header)
	ctx := util.NewContextFromGinContext(c)
	webCommon.SetProxyHeader(c)
	header := c.Request.Header
	defLang := s.Language.CreateDefaultCCLanguageIf(util.GetLanguage(header))
	defErr := s.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(header))
	customFieldsStr := c.PostForm(common.ExportCustomFields)

	hostIDStr := c.PostForm("bk_host_id")
	appIDStr := c.PostForm("bk_biz_id")
	exportCondStr := c.PostForm("export_condition")
	appID, err := strconv.ParseInt(appIDStr, 10, 64)
	if err != nil {
		blog.Errorf("ExportHost failed, bk_biz_id not integer. err: %+v, biz id: %s,  rid: %s", err, appIDStr, rid)
		err := defErr.CCErrorf(common.CCErrCommParamsNeedInt, common.BKAppIDField)
		reply := getReturnStr(err.GetCode(), err.Error(), nil)
		_, _ = c.Writer.Write([]byte(reply))
		return
	}

	objectName, objIDs, err := s.getCustomObjectInfo(ctx, header)
	if err != nil {
		blog.Errorf("get custom instance name failed, err: %v, rid: %s", err, rid)
		return
	}

	objID := common.BKInnerObjIDHost
	filterFields := logics.GetFilterFields(objID)
	customFields := logics.GetCustomFields(filterFields, customFieldsStr)
	// customLen+5为生成主机数据的起始列索引, 5=字段说明1列+业务拓扑，业务名，集群，模块4列
	fields, err := s.Logics.GetObjFieldIDs(objID, filterFields, customFields, c.Request.Header, appID, len(objectName)+5)
	if nil != err {
		blog.Errorf("get host model fields failed, err: %v, rid: %s", err, rid)
		reply := getReturnStr(common.CCErrCommExcelTemplateFailed, defErr.Errorf(common.CCErrCommExcelTemplateFailed,
			objID).Error(), nil)
		_, _ = c.Writer.Write([]byte(reply))
		return
	}

	var hostFields []string
	for _, property := range fields {
		hostFields = append(hostFields, property.ID)
	}

	hostInfo, err := s.Logics.GetHostData(appID, hostIDStr, hostFields, exportCondStr, header, defLang)
	if err != nil {
		blog.Errorf("get hosts failed, host id:%s, err: %v, rid: %s", hostIDStr, err, rid)
		reply := getReturnStr(common.CCErrWebGetHostFail, defErr.Errorf(common.CCErrWebGetHostFail,
			err.Error()).Error(), nil)
		_, _ = c.Writer.Write([]byte(reply))
		return
	}
	if len(hostInfo) == 0 {
		blog.Errorf("not find host, host id: %s, cond: %#v, rid: %s", hostIDStr, exportCondStr, rid)
		reply := getReturnStr(common.CCErrWebGetHostFail, defErr.Errorf(common.CCErrWebGetHostFail,
			"no host is found").Error(), nil)
		_, _ = c.Writer.Write([]byte(reply))
		return
	}

	err = s.handleModule(hostInfo, rid)
	if err != nil {
		blog.Errorf("add module name to host failed, err: %v, rid: %s", err, rid)
		return
	}
	setDIs, hostSetMap, err := s.handleSet(hostInfo, rid)
	if err != nil {
		blog.Errorf("add set name to host failed, err: %v, rid: %s", err, rid)
		return
	}

	if len(objIDs) > 0 {
		setParentIDs, setCustomMap, err := s.getSetParentID(ctx, header, setDIs, rid)
		if err != nil {
			blog.Errorf("get set parent id and host set rel map failed, err: %v, rid: %s", err, rid)
			return
		}

		err = s.handleCustomData(ctx, header, hostInfo, objIDs, rid, setParentIDs, setCustomMap, hostSetMap)
		if err != nil {
			blog.Errorf("get custom parent id and host custom rel map failed, err: %v, rid: %s", err, rid)
			return
		}
	}

	var file *xlsx.File
	file = xlsx.NewFile()

	usernameMap, propertyList, err := s.getUsernameMapWithPropertyList(c, objID, hostInfo)
	if nil != err {
		blog.Errorf("ExportHost failed, get username map and property list failed, err: %+v, rid: %s", err, rid)
		reply := getReturnStr(common.CCErrWebGetUsernameMapFail, defErr.Errorf(common.CCErrWebGetUsernameMapFail,
			objID).Error(), nil)
		_, _ = c.Writer.Write([]byte(reply))
		return
	}

	err = s.Logics.BuildHostExcelFromData(ctx, objID, fields, nil, hostInfo, file, header, appID, usernameMap,
		propertyList, objectName, objIDs)
	if nil != err {
		blog.Errorf("ExportHost failed, BuildHostExcelFromData failed, object:%s, err:%+v, rid:%s", objID, err,
			rid)
		reply := getReturnStr(common.CCErrCommExcelTemplateFailed, defErr.Errorf(common.CCErrCommExcelTemplateFailed,
			objID).Error(), nil)
		_, _ = c.Writer.Write([]byte(reply))
		return
	}

	dirFileName := fmt.Sprintf("%s/export", webCommon.ResourcePath)
	_, err = os.Stat(dirFileName)
	if nil != err {
		if err := os.MkdirAll(dirFileName, os.ModeDir|os.ModePerm); err != nil {
			blog.Errorf("ExportHost failed, make local dir to save export file failed, err: %+v, rid: %s", err, rid)
			c.String(http.StatusInternalServerError, fmt.Sprintf("make local dir to save export file failed, err: %+v", err))
			return
		}
	}
	fileName := fmt.Sprintf("%dhost.xlsx", time.Now().UnixNano())
	dirFileName = fmt.Sprintf("%s/%s", dirFileName, fileName)

	logics.ProductExcelCommentSheet(ctx, file, defLang)
	err = file.Save(dirFileName)
	if err != nil {
		blog.Errorf("ExportHost failed, save file failed, err: %+v, rid: %s", err, rid)
		reply := getReturnStr(common.CCErrWebCreateEXCELFail, defErr.Errorf(common.CCErrCommExcelTemplateFailed,
			err.Error()).Error(), nil)
		_, _ = c.Writer.Write([]byte(reply))
		return
	}
	logics.AddDownExcelHttpHeader(c, "bk_cmdb_export_host.xlsx")
	c.File(dirFileName)

	if err := os.Remove(dirFileName); err != nil {
		blog.Errorf("ExportHost success, but remove host.xlsx file failed, err: %+v, rid: %s", err, rid)
	}
}

// BuildDownLoadExcelTemplate build download excel template
func (s *Service) BuildDownLoadExcelTemplate(c *gin.Context) {
	rid := util.GetHTTPCCRequestID(c.Request.Header)
	ctx := util.NewContextFromGinContext(c)

	webCommon.SetProxyHeader(c)
	objID := c.Param(common.BKObjIDField)
	randNum := rand.Uint32()
	dir := webCommon.ResourcePath + "/template/"
	_, err := os.Stat(dir)
	if nil != err {
		if err := os.MkdirAll(dir, os.ModeDir|os.ModePerm); err != nil {
			blog.Errorf("BuildDownLoadExcelTemplate failed, make template dir failed, err: %+v, rid: %s", err, rid)
			c.String(http.StatusInternalServerError, fmt.Sprintf("make template dir failed, err: %+v", err))
			return
		}
	}
	language := webCommon.GetLanguageByHTTPRequest(c)
	defLang := s.Language.CreateDefaultCCLanguageIf(language)
	defErr := s.CCErr.CreateDefaultCCErrorIf(language)

	modelBizID, err := parseModelBizID(c.PostForm(common.BKAppIDField))
	if err != nil {
		msg := getReturnStr(common.CCErrCommJSONUnmarshalFailed, defErr.Error(common.CCErrCommJSONUnmarshalFailed).Error(), nil)
		c.String(http.StatusOK, msg)
		return
	}

	file := fmt.Sprintf("%s/%stemplate-%d-%d.xlsx", dir, objID, time.Now().UnixNano(), randNum)
	err = s.Logics.BuildExcelTemplate(ctx, objID, file, c.Request.Header, defLang, modelBizID)
	if nil != err {
		blog.Errorf("BuildDownLoadExcelTemplate failed, build excel template failed, object:%s error:%s, rid: %s", objID, err.Error(), rid)
		reply := getReturnStr(common.CCErrCommExcelTemplateFailed, defErr.Errorf(common.CCErrCommExcelTemplateFailed, objID).Error(), nil)
		_, _ = c.Writer.Write([]byte(reply))
		return
	}
	if objID == common.BKInnerObjIDHost {
		logics.AddDownExcelHttpHeader(c, "bk_cmdb_import_host.xlsx")
	} else {
		logics.AddDownExcelHttpHeader(c, fmt.Sprintf("bk_cmdb_inst_%s.xlsx", objID))
	}

	http.ServeFile(c.Writer, c.Request, file)
	c.File(file)
	if err := os.Remove(file); err != nil {
		blog.Errorf("BuildDownLoadExcelTemplate success, but remove template file after response failed, err: %+v, rid: %s", err, rid)
	}
	return
}

// getReturnStr get return string
func getReturnStr(code int, message string, data interface{}) string {
	ret := make(map[string]interface{})
	ret["bk_error_code"] = code
	if 0 == code {
		ret["result"] = true
	} else {
		ret["result"] = false
	}
	ret["bk_error_msg"] = message
	ret["data"] = data
	msg, _ := json.Marshal(ret)

	return string(msg)

}

func (s *Service) ListenIPOptions(c *gin.Context) {
	rid := util.GetHTTPCCRequestID(c.Request.Header)
	ctx := util.NewContextFromGinContext(c)
	webCommon.SetProxyHeader(c)
	header := c.Request.Header
	defErr := s.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(header))

	hostIDStr := c.Param("bk_host_id")
	hostID, err := strconv.ParseInt(hostIDStr, 10, 64)
	if err != nil {
		blog.Infof("host id invalid, convert to int failed, hostID: %s, err: %+v, rid: %s", hostID, err, rid)
		result := metadata.ResponseDataMapStr{
			BaseResp: metadata.BaseResp{
				Result: false,
				Code:   common.CCErrCommParamsInvalid,
				ErrMsg: defErr.Errorf(common.CCErrCommParamsInvalid, common.BKHostIDField).Error(),
			},
		}
		c.JSON(http.StatusOK, result)
		return
	}
	option := metadata.ListHostsWithNoBizParameter{
		HostPropertyFilter: &querybuilder.QueryFilter{
			Rule: querybuilder.CombinedRule{
				Condition: querybuilder.ConditionAnd,
				Rules: []querybuilder.Rule{
					querybuilder.AtomRule{
						Field:    common.BKHostIDField,
						Operator: querybuilder.OperatorEqual,
						Value:    hostID,
					},
				},
			},
		},
		Fields: []string{
			common.BKHostIDField,
			common.BKHostNameField,
			common.BKHostInnerIPField,
			common.BKHostOuterIPField,
		},
		Page: metadata.BasePage{
			Start: 0,
			Limit: 1,
		},
	}
	resp, err := s.CoreAPI.ApiServer().ListHostWithoutApp(ctx, c.Request.Header, option)
	if err != nil {
		blog.Errorf("get host by id failed, hostID: %d, err: %+v, rid: %s", hostID, err, rid)
		result := metadata.ResponseDataMapStr{
			BaseResp: metadata.BaseResp{
				Result: false,
				Code:   common.CCErrHostGetFail,
				ErrMsg: defErr.Error(common.CCErrHostGetFail).Error(),
			},
		}
		c.JSON(http.StatusOK, result)
		return
	}
	if resp.Code != 0 || resp.Result == false {
		blog.Errorf("got host by id failed, hostID: %d, response: %+v, rid: %s", hostID, resp, rid)
		c.JSON(http.StatusOK, resp)
		return
	}
	if len(resp.Data.Info) == 0 {
		blog.Errorf("host not found, hostID: %d, rid: %s", hostID, rid)
		result := metadata.ResponseDataMapStr{
			BaseResp: metadata.BaseResp{
				Result: false,
				Code:   common.CCErrCommNotFound,
				ErrMsg: defErr.Error(common.CCErrCommNotFound).Error(),
			},
		}
		c.JSON(http.StatusOK, result)
		return
	}
	type Host struct {
		HostID   int64  `json:"bk_host_id" bson:"bk_host_id"`           // 主机ID(host_id)								数字
		HostName string `json:"bk_host_name" bson:"bk_host_name"`       // 主机名称
		InnerIP  string `json:"bk_host_innerip" bson:"bk_host_innerip"` // 内网IP
		OuterIP  string `json:"bk_host_outerip" bson:"bk_host_outerip"` // 外网IP
	}
	host := Host{}
	raw := resp.Data.Info[0]
	if err := mapstr.DecodeFromMapStr(&host, raw); err != nil {
		msg := fmt.Sprintf("decode response data into host failed, raw: %+v, err: %+v, rid: %s", raw, err, rid)
		blog.Error(msg)
		result := metadata.ResponseDataMapStr{
			BaseResp: metadata.BaseResp{
				Result: false,
				Code:   common.CCErrCommJSONUnmarshalFailed,
				ErrMsg: defErr.Error(common.CCErrCommJSONUnmarshalFailed).Error(),
			},
		}
		c.JSON(http.StatusOK, result)
		return
	}

	ipOptions := make([]string, 0)
	ipOptions = append(ipOptions, "127.0.0.1")
	ipOptions = append(ipOptions, "0.0.0.0")
	if len(host.InnerIP) > 0 {
		ipOptions = append(ipOptions, host.InnerIP)
	}
	if len(host.OuterIP) > 0 {
		ipOptions = append(ipOptions, host.OuterIP)
	}
	result := metadata.ResponseDataMapStr{
		BaseResp: metadata.BaseResp{
			Result: true,
			Code:   0,
		},
		Data: map[string]interface{}{
			"options": ipOptions,
		},
	}
	c.JSON(http.StatusOK, result)
	return
}

// UpdateHosts Excel update host batch
func (s *Service) UpdateHosts(c *gin.Context) {
	rid := util.GetHTTPCCRequestID(c.Request.Header)
	ctx := util.NewContextFromHTTPHeader(c.Request.Header)

	language := webCommon.GetLanguageByHTTPRequest(c)
	defLang := s.Language.CreateDefaultCCLanguageIf(language)
	defErr := s.CCErr.CreateDefaultCCErrorIf(language)
	bizIDStr := c.PostForm("bk_biz_id")
	bizID, err := strconv.ParseInt(bizIDStr, 10, 64)
	if err != nil {
		blog.Errorf("UpdateHosts failed, bk_biz_id not integer. err: %+v, biz id: %s,  rid: %s", err, bizIDStr, rid)
		err := defErr.CCErrorf(common.CCErrCommParamsNeedInt, common.BKAppIDField)
		reply := getReturnStr(err.GetCode(), err.Error(), nil)
		_, _ = c.Writer.Write([]byte(reply))
		return
	}
	file, err := c.FormFile("file")
	if nil != err {
		blog.Errorf("UpdateHost excel import update hosts failed, get file from form data failed, err: %+v, rid: %s", err, rid)
		msg := getReturnStr(common.CCErrWebFileNoFound, defErr.Error(common.CCErrWebFileNoFound).Error(), nil)
		c.String(http.StatusOK, string(msg))
		return
	}
	webCommon.SetProxyHeader(c)

	randNum := rand.Uint32()
	dir := webCommon.ResourcePath + "/import/"
	_, err = os.Stat(dir)
	if nil != err {
		if err := os.MkdirAll(dir, os.ModeDir|os.ModePerm); err != nil {
			blog.Errorf("UpdateHost excel import update hosts, save form data to local file failed, mkdir failed, err: %+v, rid: %s", err, rid)
			c.String(http.StatusInternalServerError, fmt.Sprintf("save form data to local file failed, mkdir failed, err: %+v", err))
			return
		}
	}
	filePath := fmt.Sprintf("%s/importhost-%d-%d.xlsx", dir, time.Now().UnixNano(), randNum)
	if err := c.SaveUploadedFile(file, filePath); nil != err {
		blog.Errorf("UpdateHosts failed, save form data to local file failed, save data as excel failed, err: %+v, rid: %s", err, rid)
		msg := getReturnStr(common.CCErrWebFileSaveFail, defErr.Errorf(common.CCErrWebFileSaveFail, err.Error()).Error(), nil)
		c.String(http.StatusOK, string(msg))
		return
	}

	// del file
	defer func(filePath string, rid string) {
		if err := os.Remove(filePath); err != nil {
			blog.Errorf("UpdateHost excel import update hosts, remove temporary file failed, err: %+v, rid: %s", err, rid)
		}
	}(filePath, rid)

	f, err := xlsx.OpenFile(filePath)
	if nil != err {
		blog.Errorf("UpdateHost excel import update hosts failed, open form data as excel file failed, err: %+v, rid: %s", err, rid)
		msg := getReturnStr(common.CCErrWebOpenFileFail, defErr.Errorf(common.CCErrWebOpenFileFail, err.Error()).Error(), nil)
		c.String(http.StatusOK, string(msg))
		return
	}
	result := s.Logics.UpdateHosts(ctx, f, c.Request.Header, defLang, bizID)

	c.JSON(http.StatusOK, result)
}

func (s *Service) getCustomData(ctx context.Context, header http.Header, instIDs []int64, objID, rid string) ([]int64,
	map[int64]int64, map[int64]string, error) {
	query := &metadata.QueryCondition{
		Condition: mapstr.MapStr{
			common.BKInstIDField: mapstr.MapStr{
				common.BKDBIN: instIDs,
			},
		},
		Fields: []string{common.BKInstIDField, common.BKInstNameField, common.BKInstParentStr},
	}

	insts, err := s.Engine.CoreAPI.ApiServer().ReadInstance(ctx, header, objID, query)
	if err != nil {
		blog.Errorf("get custom level inst data failed, query cond: %#v, err: %v, rid: %s", query, err, rid)
		return nil, nil, nil, err
	}

	parentIDs := make([]int64, 0)
	instIdParentIdMap := make(map[int64]int64, 0)
	instIdNameMap := make(map[int64]string, 0)
	for _, inst := range insts.Data.Info {
		parentID, err := inst.Int64(common.BKParentIDField)
		if err != nil {
			blog.Errorf("get inst parent id failed, err: %v, rid: %s", err, rid)
			return nil, nil, nil, err
		}
		parentIDs = append(parentIDs, parentID)

		instID, err := inst.Int64(common.BKInstIDField)
		if err != nil {
			blog.Errorf("get inst id failed, err: %v, rid: %s", err, rid)
			return nil, nil, nil, err
		}
		instIdParentIdMap[instID] = parentID

		instName, err := inst.String(common.BKInstNameField)
		if err != nil {
			blog.Errorf("get inst name failed, err: %v, rid: %s", err, rid)
			return nil, nil, nil, err
		}
		instIdNameMap[instID] = instName
	}

	return parentIDs, instIdParentIdMap, instIdNameMap, nil
}

// getCustomObjectInfo get custom instance object info
func (s *Service) getCustomObjectInfo(ctx context.Context, header http.Header) ([]string, []string, error) {
	rid := util.ExtractRequestIDFromContext(ctx)
	query := &metadata.QueryCondition{
		Condition: mapstr.MapStr{
			common.AssociationKindIDField: common.AssociationKindMainline,
		},
	}
	mainlineAsstRsp, err := s.Engine.CoreAPI.ApiServer().ReadModuleAssociation(context.Background(), header, query)
	if err != nil {
		blog.Errorf("search mainline association failed, err: %v, rid: %s", err, rid)
		return nil, nil, err
	}

	mainlineObjectChildMap := make(map[string]string, 0)
	for _, asst := range mainlineAsstRsp.Data.Info {
		if asst.ObjectID == common.BKInnerObjIDHost {
			continue
		}
		mainlineObjectChildMap[asst.AsstObjID] = asst.ObjectID
	}

	// get all mainline custom object id
	objectIDs := make([]string, 0)
	for objectID := common.BKInnerObjIDApp; len(objectID) != 0; objectID = mainlineObjectChildMap[objectID] {
		if objectID == common.BKInnerObjIDApp || objectID == common.BKInnerObjIDSet ||
			objectID == common.BKInnerObjIDModule {
			continue
		}
		objectIDs = append(objectIDs, objectID)
	}
	if len(objectIDs) == 0 {
		return nil, nil, nil
	}

	input := &metadata.QueryCondition{
		Fields: []string{common.BKObjNameField, common.BKObjIDField},
		Condition: mapstr.MapStr{
			common.BKObjIDField: mapstr.MapStr{common.BKDBIN: objectIDs},
		},
	}

	objectName := make([]string, 0)
	objects, err := s.Logics.CoreAPI.ApiServer().ReadModel(context.Background(), header, input)
	if err != nil {
		blog.Errorf("search mainline obj failed, objIDs: %#v, err: %v, rid: %s", objectIDs, err, rid)
		return objectName, util.ReverseArrayString(objectIDs), err
	}

	objIDNameMap := make(map[string]string, 0)
	for _, val := range objects.Data.Info {
		objIDNameMap[val.Spec.ObjectID] = val.Spec.ObjectName
	}

	for _, objID := range objectIDs {
		if objName, ok := objIDNameMap[objID]; ok {
			objectName = append(objectName, objName)
		}
	}

	return objectName, util.ReverseArrayString(objectIDs), nil
}

// handleModule 处理module数据
func (s *Service) handleModule(hostInfo []mapstr.MapStr, rid string) error {
	// 统计host与module关系
	for _, data := range hostInfo {
		moduleMap, exist := data[common.BKInnerObjIDModule].([]interface{})
		if !exist {
			blog.Errorf("get module map data from host data failed, not exist, data: %#v, rid: %s", data, rid)
			return fmt.Errorf("from host data get module map, not exist, rid: %s", rid)
		}

		moduleNameMap := make(map[string]int)
		for idx, module := range moduleMap {
			rowMap, err := mapstr.NewFromInterface(module)
			if err != nil {
				blog.Errorf("get module data from host data failed, err: %v, rid: %s", err, rid)
				return err
			}

			moduleName, err := rowMap.String(common.BKModuleNameField)
			if err != nil {
				blog.Errorf("get module name from host data failed, err: %v, rid: %s", err, rid)
				return fmt.Errorf("from host data get module name, not exist, rid: %s", rid)
			}
			moduleNameMap[moduleName] = idx
		}

		var moduleStr string
		for moduleName := range moduleNameMap {
			if moduleStr == "" {
				moduleStr = moduleName
			} else {
				moduleStr += "," + moduleName
			}
		}
		data.Set("modules", moduleStr)
	}

	return nil
}

// handleModule 处理set数据
func (s *Service) handleSet(hostInfo []mapstr.MapStr, rid string) ([]int64, map[int64][]int64, error) {
	// 统计host与set关系
	hostSetMap := make(map[int64][]int64, 0)
	setIDs := make([]int64, 0)
	header := util.BuildHeader(common.CCSystemOperatorUserName, common.BKDefaultOwnerID)
	res, err := s.CoreAPI.CoreService().System().SearchPlatformSetting(context.Background(), header)
	if err != nil {
		return nil, nil, err
	}
	conf := res.Data

	for _, data := range hostInfo {
		setMap, exist := data[common.BKInnerObjIDSet].([]interface{})
		if !exist {
			blog.Errorf("get set map data from host data, not exist, data: %#v, rid: %s", data, rid)
			return nil, nil, fmt.Errorf("from host data get set map, not exist, rid: %s", rid)
		}

		rowMap, err := mapstr.NewFromInterface(data[common.BKInnerObjIDHost])
		if err != nil {
			blog.Errorf("get host map data failed, hostData: %#v, err: %v, rid: %s", data, err, rid)
			return nil, nil, err
		}

		hostID, err := rowMap.Int64(common.BKHostIDField)
		if err != nil {
			blog.Errorf("get host id failed, host id: %s, err: %v, rid: %s", hostID, err, rid)
			return nil, nil, nil
		}

		setNameMap := make(map[string]int)
		setSubIDs := make([]int64, 0)
		for idx, set := range setMap {
			rowMap, err := mapstr.NewFromInterface(set)
			if err != nil {
				blog.Errorf("get set data from host data failed, err: %v, rid: %s", err, rid)
				return nil, nil, err
			}

			setName, err := rowMap.String(common.BKSetNameField)
			if err != nil {
				blog.Errorf("get set name from host data failed, err: %v, rid: %s", err, rid)
				return nil, nil, fmt.Errorf("from host data get set name, not exist, rid: %s", rid)
			}
			setNameMap[setName] = idx

			setID, err := rowMap.Int64(common.BKSetIDField)
			if err != nil {
				blog.Errorf("get set id from host data failed, err: %v, rid: %s", err, rid)
				return nil, nil, err
			}

			if setName != string(conf.BuiltInSetName) {
				setIDs = append(setIDs, setID)
				setSubIDs = append(setSubIDs, setID)
			}
		}

		hostSetMap[hostID] = setSubIDs

		var setStr string
		for setName := range setNameMap {
			if setStr == "" {
				setStr = setName
			} else {
				setStr += "," + setName
			}
		}
		data.Set("sets", setStr)
	}

	return setIDs, hostSetMap, nil
}

// getSetParentID get set parent id and set custom rel map
func (s *Service) getSetParentID(ctx context.Context, header http.Header, setIDs []int64, rid string) ([]int64,
	map[int64]int64, error) {
	// 获取set数据，统计set parent id
	querySet := &metadata.QueryCondition{
		Condition: mapstr.MapStr{
			common.BKSetIDField: mapstr.MapStr{
				common.BKDBIN: setIDs,
			},
		},
		Fields: []string{common.BKSetIDField, common.BKInstParentStr, common.BKSetNameField},
	}

	sets, err := s.Engine.CoreAPI.ApiServer().ReadInstance(ctx, header, common.BKInnerObjIDSet, querySet)
	if err != nil {
		blog.Errorf("get set data failed, cond: %#v, err: %v,rid:%s", querySet, err, rid)
		return nil, nil, err
	}
	if !sets.Result {
		blog.Errorf("get sets failed, err code: %d, err msg: %s, rid: %s", sets.Code, sets.ErrMsg, rid)
		return nil, nil, fmt.Errorf("get sets failed, err msg: %s", sets.ErrMsg)
	}

	setParentIDs := make([]int64, 0)
	setCustomMap := make(map[int64]int64, 0)
	for _, set := range sets.Data.Info {
		parentID, err := set.Int64(common.BKInstParentStr)
		if err != nil {
			blog.Errorf("get set parent id failed, err: %v, rid: %s", err, rid)
			return nil, nil, err
		}
		setParentIDs = append(setParentIDs, parentID)

		setID, err := set.Int64(common.BKSetIDField)
		if err != nil {
			blog.Errorf("get set id failed, err: %v, rid: %s", err, rid)
			return nil, nil, err
		}
		setCustomMap[setID] = parentID
	}

	return setParentIDs, setCustomMap, nil
}

// handleCustomData 处理自定义成层级数据
func (s *Service) handleCustomData(ctx context.Context, header http.Header, hostInfo []mapstr.MapStr, objIDs []string,
	rid string, parentIDs []int64, setCustomMap map[int64]int64, hostSetMap map[int64][]int64) error {
	instIdParentIDMap := make(map[int64]int64, 0)
	instIdNameMap := make(map[int64]string, 0)
	var err error
	for _, objID := range objIDs {
		parentIDs, instIdParentIDMap, instIdNameMap, err = s.getCustomData(ctx, header, parentIDs, objID, rid)
		if err != nil {
			blog.Errorf("get custom data failed, cond: %#v, err: %v, rid: %s", parentIDs, err, rid)
			return err
		}

		hostCustomMap := make(map[int64][]int64, 0)
		hostCustomNameMap := make(map[int64]string, 0)
		for hostID, setIDs := range hostSetMap {
			customNameMap := make(map[string]int, 0)
			for idx, setID := range setIDs {
				customNameMap[instIdNameMap[setCustomMap[setID]]] = idx
				hostCustomMap[hostID] = append(hostCustomMap[hostID], setCustomMap[setID])
			}

			customStr := ""
			for customName := range customNameMap {
				if customStr == "" {
					customStr = customName
				} else {
					customStr += "," + customName
				}
			}

			hostCustomNameMap[hostID] = customStr
		}

		for _, data := range hostInfo {
			rowMap, err := mapstr.NewFromInterface(data[common.BKInnerObjIDHost])
			if err != nil {
				blog.Errorf("get host map data failed, hostData: %#v, err: %v, rid: %s", data, err, rid)
				return err
			}

			hostID, err := rowMap.Int64(common.BKHostIDField)
			if err != nil {
				blog.Errorf("get host id failed, host id: %s, err: %v, rid: %s", hostID, err, rid)
				return err
			}

			data[objID] = hostCustomNameMap[hostID]
		}

		setCustomMap = instIdParentIDMap
		hostSetMap = hostCustomMap
	}

	return nil
}
