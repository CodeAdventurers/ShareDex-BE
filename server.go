package main

import (
	"gin_demo/routers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// 设置 common 路由
	routers.SetupCommonRoutes(r)

	return r
}
func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
