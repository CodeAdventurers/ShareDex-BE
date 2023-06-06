package routers

import (
	"ShareDex-BE/src/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRouter struct{}

// userRouters 用户模块路由
var userRouters = []RouterInfo{
	{
		// 获取用户信息
		Method: http.MethodGet, Path: "/user_service/hui", Handler: handler.GetUserInfo,
		//ReqIn: api_models.UserIn{}, RespOut: api_models.UserOut{}, // 方便快速知道请求入参和出参
	},

	{
		// 注册用户信息
		Method: http.MethodGet, Path: "/user_service/register", Handler: handler.GetUserInfo,
		//ReqIn: api_models.UserIn{}, RespOut: api_models.UserOut{}, // 方便快速知道请求入参和出参
	},
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	for _, routerInfo := range userRouters {
		router.Handle(routerInfo.Method, routerInfo.Path, routerInfo.Handler)
	}
}
