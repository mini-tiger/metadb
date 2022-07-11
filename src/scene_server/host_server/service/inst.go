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
	"configcenter/src/common/auditlog"
	"strconv"

	"configcenter/src/ac/iam"
	"configcenter/src/common"
	"configcenter/src/common/auth"
	"configcenter/src/common/blog"
	"configcenter/src/common/errors"
	"configcenter/src/common/http/rest"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	meta "configcenter/src/common/metadata"
	"configcenter/src/common/util"
)

// CreatePlatBatch create plat instance in batch
func (s *Service) CreateManyBatch(ctx *rest.Contexts) {
	objID := ctx.Request.PathParameter("bk_obj_id")
	input := struct {
		Data []mapstr.MapStr `json:"data"`
	}{}

	if err := ctx.DecodeInto(&input); nil != err {
		ctx.RespAutoError(err)
		return
	}

	if len(input.Data) == 0 {
		blog.Errorf("CreateMany , input is empty, rid:%s", ctx.Kit.Rid)
		ctx.RespAutoError(ctx.Kit.CCError.CCError(common.CCErrCommHTTPBodyEmpty))
		return
	}

	user := util.GetUser(ctx.Request.Request.Header)
	//for i := range input.Data {
	//	input.Data[i][common.BKCreator] = user
	//	input.Data[i][common.BKLastEditor] = user
	//}

	instInfo := &meta.CreateManyModelInstance{
		Datas: input.Data,
	}

	result := metadata.CreateManyDataResult{}
	txnErr := s.Engine.CoreAPI.CoreService().Txn().AutoRunTxn(ctx.Kit.Ctx, ctx.Kit.Header, func() error {
		var err error
		res, err := s.CoreAPI.CoreService().Instance().CreateManyInstance(ctx.Kit.Ctx, ctx.Kit.Header, objID, instInfo)
		if nil != err {
			blog.Errorf("CreateMany failed, CreateManyInstance error: %s, input:%+v,rid:%s", err.Error(), input, ctx.Kit.Rid)
			return ctx.Kit.CCError.CCError(common.CCErrTopoInstCreateFailed)
		}

		if false == res.Result {
			blog.Errorf("CreateMany failed, CreateManyInstance error.err code:%d,err msg:%s,input:%+v,rid:%s", res.Code, res.ErrMsg, input, ctx.Kit.Rid)
			return errors.New(res.Code, res.ErrMsg)
		}

		if len(res.Data.Exceptions) > 0 {
			blog.Errorf("CreateMany failed, err:#v,input:%+v,rid:%s", res.Data.Exceptions, input, ctx.Kit.Rid)
			return ctx.Kit.CCError.New(int(res.Data.Exceptions[0].Code), res.Data.Exceptions[0].Message)
		}

		if len(res.Data.Created) == 0 {
			blog.Errorf("CreateMany failed, no obj was found,input:%+v,rid:%s", input, ctx.Kit.Rid)
			return ctx.Kit.CCError.CCError(common.CCErrTopoCloudNotFound)
		}

		//platIDs := make([]int64, len(res.Data.Created))
		//for i, created := range res.Data.Created {
		//	platIDs[i] = int64(created.ID)
		//	result[i] = mapstr.MapStr{
		//		"ID": int64(created.ID),
		//	}
		//}
		result = res.Data

		// generate audit log.
		//audit := auditlog.NewCloudAreaAuditLog(s.CoreAPI.CoreService())
		audit := auditlog.NewInstanceAudit(s.CoreAPI.CoreService())
		generateAuditParameter := auditlog.NewGenerateAuditCommonParameter(ctx.Kit, metadata.AuditCreate)

		auditdata := make([]mapstr.MapStr, len(result.Created))
		//generateAuditParameter := auditlog.NewGenerateAuditCommonParameter(kit, metadata.AuditCreate)
		//auditLog, err := audit.GenerateAuditLogByCondGetData(generateAuditParameter, obj.GetObjectID(), cond)
		for index, value := range result.Created {
			auditdata[index] = mapstr.MapStr{common.BKInstIDField: value.ID}
		}
		//for _, value := range result.Repeated {
		//	auditdata = append(auditdata, mapstr.MapStr{common.BKInstIDField: value.Data})
		//}

		logs, err := audit.GenerateAuditLog(generateAuditParameter, objID, auditdata)
		//if err != nil {
		//	blog.Errorf("generate audit log failed after create cloud area, err: %v, rid: %s", err, ctx.Kit.Rid)
		//	return err
		//}

		// save audit log.
		if err := audit.SaveAuditLog(ctx.Kit, logs...); err != nil {
			blog.Errorf("save audit log failed after create cloud area, err: %v, rid: %s", err, ctx.Kit.Rid)
			return err
		}

		// register cloud area resource creator action to iam
		if auth.EnableAuthorize() {
			iamInstances := make([]metadata.IamInstance, len(res.Data.Created))
			for index, created := range res.Data.Created {
				iamInstances[index] = metadata.IamInstance{
					ID:   strconv.FormatUint(created.ID, 10),
					Name: util.GetStrByInterface(input.Data[created.OriginIndex][common.BKCloudNameField]),
				}
			}
			iamInstancesWithCreator := metadata.IamInstancesWithCreator{
				Type:      string(iam.SysCloudArea),
				Instances: iamInstances,
				Creator:   user,
			}
			_, err = s.AuthManager.Authorizer.BatchRegisterResourceCreatorAction(ctx.Kit.Ctx, ctx.Kit.Header, iamInstancesWithCreator)
			if err != nil {
				blog.Errorf("register created cloud area to iam failed, err: %s, rid: %s", err, ctx.Kit.Rid)
				return err
			}
		}

		return nil
	})

	if txnErr != nil {
		ctx.RespAutoError(txnErr)
		return
	}

	ctx.RespEntity(result)

}
