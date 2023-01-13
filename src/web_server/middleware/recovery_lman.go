package middleware

import (
	"configcenter/src/web_server/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http/httputil"
)

/**
 * @Author: Tao Jun
 * @Since: 2023/1/4
 * @Desc: recovery_extend.go
**/
func RecoverExtend(out io.Writer) gin.HandlerFunc {

	var logger *log.Logger

	logger = log.New(out, "\n\n\x1b[31m", log.LstdFlags)

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if logger != nil {
					httprequest, _ := httputil.DumpRequest(c.Request, false)
					logger.Printf("[Recovery] panic recovered:\n%s\n%s\n", string(httprequest), err)
				}
				common.ResponseError(c, common.CreateRespErrStr(fmt.Sprintf("[Recovery] panic recovered : %s", err)))
				//c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}
