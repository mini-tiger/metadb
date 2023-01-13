package common

import (
	"configcenter/src/common/metadata"
	"configcenter/src/web_server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: Tao Jun
 * @Since: 2023/1/3
 * @Desc: http.go
**/

func CreateRespErr(err error) *models.BaseErr {
	return &models.BaseErr{
		Code:    "",
		Message: err.Error(),
	}
}

func CreateRespErrStr(err string) *models.BaseErr {
	return &models.BaseErr{
		Code:    "",
		Message: err,
	}
}

func ResponseBaseSuccess(c *gin.Context, reqid string, data interface{}) {
	c.JSON(http.StatusOK, models.BaseResp{
		RequestID: reqid,
		Data:      data,
	})
}

func ResponseSuccess(c *gin.Context, reqid string, data *metadata.InstResult) {
	c.JSON(http.StatusOK, models.BaseResp{
		RequestID: reqid,
		Data:      data,
	})
}

func ResponseError(c *gin.Context, err *models.BaseErr) {
	c.JSON(http.StatusOK, models.BaseResp{
		Error:     err,
		RequestID: "",
	})
}
