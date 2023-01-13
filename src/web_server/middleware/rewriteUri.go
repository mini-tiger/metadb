package middleware

import (
	"bytes"
	"configcenter/src/apimachinery/discovery"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

/**
 * @Author: Tao Jun
 * @Since: 2023/1/11
 * @Desc: rewriteUri.go
**/

// bodyLogWriter 定义一个存储响应内容的结构体
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write 读取响应数据
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ReWriteReqUri(disc discovery.DiscoveryInterface) gin.HandlerFunc {

	return func(c *gin.Context) {
		//rid := util.GetHTTPCCRequestID(c.Request.Header)
		//fmt.Println(c.Request.RequestURI)
		//fmt.Println(c.Request.URL.Path)

		if strings.Contains(c.Request.RequestURI, "lman-cmdb") {
			c.Request.RequestURI = strings.ReplaceAll(c.Request.RequestURI, "lman-cmdb/v1", "api/v3")
			c.Request.URL.Path = strings.ReplaceAll(c.Request.RequestURI, "lman-cmdb/v1", "api/v3")
			log.Printf("Rewrite Url:%s\n", c.Request.RequestURI)
			c.Set("lman", true)
		}

		//fmt.Println(c.Request.RequestURI)
		//servers, err := disc.ApiServer().GetServers()
		//blw := &bodyLogWriter{
		//	body:           bytes.NewBufferString(""),
		//	ResponseWriter: c.Writer,
		//}
		//c.Writer = blw

		c.Next()

		//originBytes := blw.body
		//fmt.Printf("%s", originBytes)
		//// clear Origin Buffer
		//blw.body = &bytes.Buffer{}
		//blw.Write([]byte("replace data"))
		//blw.ResponseWriter.Write(blw.body.Bytes())
		//c.JSON(http.StatusOK, gin.H{"abc": 1})
		//return
	}
}
