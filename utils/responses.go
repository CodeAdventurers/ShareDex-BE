package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponseOptions struct {
	Code    int
	Message string
	Data    interface{}
}

func SuccessJsonResponse(c *gin.Context, data interface{}) {
	JsonResponse(c, http.StatusOK, JsonResponseOptions{
		Code:    0,
		Message: "ok",
		Data:    data,
	})
}

func JsonResponse(c *gin.Context, statusCode int, options JsonResponseOptions) {
	c.JSON(statusCode, gin.H{
		"code":    options.Code,
		"message": options.Message,
		"data":    options.Data,
	})
}
