package common

import (
	"gin_demo/utils"
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	utils.SuccessJsonResponse(c, nil)
}
