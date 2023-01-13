package middleware

import (
	"configcenter/src/common/blog"
	"configcenter/src/common/util"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Tao Jun
 * @Since: 2023/1/4
 * @Desc: logg.go
**/

func Logmiddleware() gin.HandlerFunc {

	return func(req *gin.Context) {

		//raw := c.Request.URL.RawQuery
		// Process request

		//end:=time.Now()
		//fmt.Println(end)
		header := req.Request.Header
		body, _ := util.PeekRequest(req.Request)
		//blog.Infof("code: %s, user: %s, rip: %s, uri: %s, body: %s, rid: %s",
		//	header.Get("Bk-App-Code"), header.Get("Bk_user"), header.Get("X-Real-Ip"),
		//	req.Request.RequestURI, body, util.GetHTTPCCRequestID(header))
		//fmt.Printf("%+v\n", req.Request)
		var rip string = header.Get("X-Real-Ip")
		if rip == "" {
			rip = req.Request.RemoteAddr
		}
		blog.Infof("Access :[%s] Host:[%s] method: [%s], rip: [%s], user: [%s],http_blueking_supplier_id: %s  body: %s, rid: %s",
			req.Request.RequestURI, req.Request.Host, req.Request.Method, rip, header.Get("Bk_user"),
			header.Get("HTTP_BLUEKING_SUPPLIER_ID"), body, util.GetHTTPCCRequestID(header))

		// xxx c.Next以上是请求时的运行
		req.Next() // 注意 next()方法的作用是跳过该调用链去直接后面的中间件以及api路由
		// xxx c.Next以下是返回时的运行

	}
}
