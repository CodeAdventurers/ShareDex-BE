package routers

import (
	"ShareDex-BE/src/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonRouter struct{}

// commonRouters 公共模块路由
var commonRouters = []RouterInfo{
	{
		// 心跳
		Method: http.MethodGet, Path: "/common/ping", Handler: handler.Ping,
	},

	{
		// 上传
		Method: http.MethodGet, Path: "/common/upload", Handler: handler.Upload,
	},
}

func (cr *CommonRouter) InitCommonRouter(router *gin.RouterGroup) {
	for _, routerInfo := range commonRouters {
		router.Handle(routerInfo.Method, routerInfo.Path, routerInfo.Handler)
	}
}
