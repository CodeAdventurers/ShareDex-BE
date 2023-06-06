package routers

import (
	"ShareDex-BE/src/data_model/api_model"
	"ShareDex-BE/src/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRouter struct{}

// userRouters 用户模块路由
var userRouters = []RouterInfo{
	{
		// 获取用户信息
		Method: http.MethodGet, Path: "/user", Handler: handler.GetUserInfo,
		ReqIn: api_model.UserIn{}, RespOut: api_model.UserOut{}, // 方便快速知道请求入参和出参
	},

	{
		// 注册用户信息
		Method: http.MethodGet, Path: "/user/register", Handler: handler.GetUserInfo,
		ReqIn: api_model.UserIn{}, RespOut: api_model.UserOut{}, // 方便快速知道请求入参和出参
	},
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	for _, routerInfo := range userRouters {
		router.Handle(routerInfo.Method, routerInfo.Path, routerInfo.Handler)
	}
}
