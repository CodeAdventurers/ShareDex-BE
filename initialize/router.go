package initialize

import (
	"github.com/CodeAdventurers/ShareDex-BE/middlewares"
	"github.com/CodeAdventurers/ShareDex-BE/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET(
		"/health", func(c *gin.Context) {
			c.JSON(
				http.StatusOK, gin.H{
					"code":    http.StatusOK,
					"success": true,
				},
			)
		},
	)

	//配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/v1/api")
	router.InitUserRouter(ApiGroup)
	return Router
}
