package service

import (
	"configcenter/src/common/json"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/topo_server/core/operation"
	"configcenter/src/web_server/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Tao Jun
 * @Since: 2023/1/3
 * @Desc: controllers.go
**/

// SearchInstByAssociation function
// @Summary 按条件查询模型实例
// @version 1.0
// @Accept  json
// @Produce  json
// @Description urlPath objectID 模型名称
// @Tags 按条件查询模型实例接口
// @Param data body models.ExtendInstFilter true "request params"
// @Success 200  object models.BaseResp "成功后返回"
// @Header  200,400,default  {string}  Token     "token"
// @Router /lman-cmdb/v3/find/instassociation/object/:objectID [get]
func (s *LmanSvc) SearchInstByAssociation(c *gin.Context) {
	//fmt.Println(c.Request.Body)
	//fmt.Println(c.Request.Header)

	reqBytes, err := c.GetRawData() // 接收json数据
	if err != nil {
		common.ResponseError(c, common.CreateRespErr(err)) //统一返回
		return
	}
	//fmt.Println(string(bytes))

	reqParse := &metadata.SearchInstResultLman{}

	objID := c.Param("obj_id")
	//fmt.Println(objID)
	err = json.Unmarshal(reqBytes, reqParse)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	//fmt.Printf("%+v\n", reqParse)
	reqSub := make([]operation.ConditionItem, 0, 1)

	reqSub = append(reqSub, operation.ConditionItem{
		Field:    reqParse.Field,
		Operator: "$in",
		Value:    reqParse.Value,
	})

	var reqProxy *operation.AssociationParams = &operation.AssociationParams{
		Condition: map[string][]operation.ConditionItem{
			objID: reqSub,
		},
	}

	url, err := s.Disc.ApiServer().GetRandomServer()
	if nil != err || url == "" {
		//blog.Errorf("no api server can be used. err: %v, rid: %s", err, rid)
		common.ResponseError(c, common.CreateRespErrStr("no api server can be used."))
		//c.Abort()
		return
	}
	//url := servers[0]
	//fmt.Println(url)
	urlReq := fmt.Sprintf("%s/api/v3/find/instassociation/object/%s", url, objID)

	respBody := &metadata.QueryInstResult{}
	header := c.Request.Header

	_, err = common.NewRP(nil, header, urlReq, reqProxy, respBody).RequestData()
	//fmt.Println(header, urlReq, reqProxy, respBody)
	if err != nil {
		common.ResponseError(c, common.CreateRespErr(err))
		return
	}

	//fmt.Printf("%+v\n", respBody)
	//fmt.Println(err)
	//fmt.Println(resp)

	common.ResponseSuccess(c, "", &respBody.Data)
	return
}
