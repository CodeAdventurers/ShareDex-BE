package routers

import (
	"gin_demo/handlers/common"
	"github.com/gin-gonic/gin"
)

func SetupCommonRoutes(r *gin.Engine) {
	commonRoutes := r.Group("/api/v1/common")
	{
		commonRoutes.GET("/ping", common.PingHandler)
	}
}
