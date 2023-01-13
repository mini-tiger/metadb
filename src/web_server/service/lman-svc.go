package service

import (
	"configcenter/src/apimachinery/discovery"
	"configcenter/src/web_server/common"
	"configcenter/src/web_server/docs"
	"configcenter/src/web_server/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

/**
 * @Author: Tao Jun
 * @Since: 2023/1/10
 * @Desc: controllers.go
**/

type LmanSvc struct {
	ExtendGin *gin.Engine
	Disc      discovery.DiscoveryInterface
	Port      string
}

func LmanServices(dis discovery.DiscoveryInterface, port string) error {
	l := &LmanSvc{
		ExtendGin: nil,
		Disc:      dis,
		Port:      port,
	}
	l.LoadExtendRoute()

	return l.Run()
}

func (s *LmanSvc) Run() error {
	return s.ExtendGin.Run(s.Port)
}

//func (s *LmanSvc) WrapHandle(e *backbone.Engine) {
//	logMiddleWare.LogMiddleware(e.Metric().HTTPMiddleware(s.ExtendGin))
//}

func (s *LmanSvc) LoadExtendRoute() *gin.Engine {
	//setGinMode()
	ws := gin.New()
	s.ExtendGin = ws
	ws.Use(gin.Logger())
	ws.Use(middleware.Logmiddleware())
	//ws.Use(gin.Recovery())

	ws.Use(middleware.RecoverExtend(os.Stderr))

	// test swagger
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	//docs.SwaggerInfo.BasePath = ""  // xxx 覆盖上面的 BasePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	ws.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//ws.POST("/issue/add", extend.IssueAdd)
	//ws.GET("/issue/list", extend.IssueList)

	// Hello World
	ws.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World,Taojun First Deploy for KubeSphere")
	})

	//v1组路由
	apiGroup := ws.Group(fmt.Sprintf("/%s/%s/", common.SuffixUrl, common.ApiVersion))

	apiGroup.GET("find/instassociation/object/:obj_id", s.SearchInstByAssociation)
	//apiGroup.Any("/*path", s.SearchInstByAssociation)

	return ws
}
