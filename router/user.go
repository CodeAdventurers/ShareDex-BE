package router

import (
	"github.com/CodeAdventurers/ShareDex-BE/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("pwd_login", api.PassWordLogin)

	}
}
