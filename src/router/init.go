package routers

import (
	"github.com/gin-gonic/gin"
)

type RouterInfo struct {
	Method  string
	Path    string
	Handler func(c *gin.Context)
	ReqIn   any // 用于知道接口的入参是什么
	RespOut any // 接口的出参
}

type AppRouterGroup struct {
	CommonRouter
	UserRouter
}

// InitRouters 初始化项目路由
func InitRouters(e *gin.Engine) {

	// api group
	apiGroup := e.Group("/api/v1")
	appRouter := new(AppRouterGroup)

	//初始化公共模块路由
	appRouter.InitCommonRouter(apiGroup)

	//初始化用户模块路由
	appRouter.InitUserRouter(apiGroup)
}
