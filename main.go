package main

import (
	"ShareDex-BE/src/dao"
	"ShareDex-BE/src/global"
	"ShareDex-BE/src/middleware"
	router "ShareDex-BE/src/router"
	"github.com/gin-gonic/gin"
)

// InitSetup 初始化项目配置
func InitSetup() {
	global.InitViper() // 	初始化viper配置

	global.InitLogger() // 初始日志器

	dao.InitMySQL() // 初始化MySQL

	dao.InitRedis() // 初始化Redis

	// ...

	global.Logger.Info("InitSetup Success")
}

func InitMiddlewares(app *gin.Engine) {

	// 日志中间件
	app.Use(middleware.ZapLogger())
}

// AppStart 项目应用启动
func AppStart() {
	app := gin.Default()

	// 初始化中间件
	InitMiddlewares(app)

	// 初始化路由
	router.InitRouters(app)

	// 程序运行
	err := app.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}

func main() {

	InitSetup()
	AppStart()
}
