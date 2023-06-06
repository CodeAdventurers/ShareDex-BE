package handler

import (
	"ShareDex-BE/src/util/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, web.SuccessResp(nil))
}

func Upload(c *gin.Context) {
	c.JSON(http.StatusOK, web.SuccessResp(nil))
}
