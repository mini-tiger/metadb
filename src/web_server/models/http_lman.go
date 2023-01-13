package models

/**
 * @Author: Tao Jun
 * @Since: 2023/1/3
 * @Desc: http.go
**/

type BaseErr struct {
	Code    string `form:"code" json:"code" example:"错误代码"`
	Message string `form:"Message" json:"Message" example:"错误描述信息"`
}

type BaseResp struct {
	Error     *BaseErr    `form:"error" json:"error,omitempty"`
	RequestID string      `form:"RequestID" json:"RequestID" example:"请求ID"`
	Data      interface{} `form:"data" json:"data" swaggertype:"object,string" example:"count:数据条数,info:数据数组对象"`
}
