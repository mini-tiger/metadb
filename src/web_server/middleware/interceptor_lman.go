package middleware

import (
	"configcenter/src/web_server/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty"
	"github.com/tidwall/gjson"
)

/**
 * @Author: Tao Jun
 * @Since: 2023/1/12
 * @Desc: lman.go
**/

var Client *resty.Client = resty.New()

func RespRebuildLman(c *gin.Context, url string) {

	urlReq := fmt.Sprintf("%s%s", url, c.Request.RequestURI)
	b, err := c.GetRawData()
	if err != nil {
		common.ResponseError(c, common.CreateRespErrStr("json Valid err"))
		return
	}

	r := Client.R()
	r.Method = c.Request.Method
	r.URL = urlReq
	r.Header = c.Request.Header

	resp, err := r.SetBody(b).Send()
	if err != nil {
		common.ResponseError(c, common.CreateRespErrStr(fmt.Sprintf("http proxy request err:%s", err.Error())))
		return
	}
	respStr := resp.String()

	if !gjson.Valid(resp.String()) {
		common.ResponseError(c, common.CreateRespErrStr("json Valid err"))
		return
	}
	respResult := gjson.Get(respStr, "result")
	if respResult.Exists() {
		if !respResult.Bool() {
			common.ResponseError(c, common.CreateRespErrStr(gjson.Get(respStr, "bk_error_msg").String()))
			return
		}
	}

	respData := gjson.Get(respStr, "data")
	if respData.Exists() {
		common.ResponseBaseSuccess(c, "", respData.Value())
		return
	}

	//switch true {
	//case strings.Contains(urlReq, "find"):
	//	result := &metadata.SearchInstResult{}
	//	resp, _ := r.SetBody(b).SetResult(result).Post(urlReq)
	//}
	//c.Data(http.StatusOK, "application/json", am.Body())
}
