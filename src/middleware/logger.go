package middleware

import (
	"ShareDex-BE/src/constant"
	"ShareDex-BE/src/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// ZapLogger 日志中间件
func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求
		c.Next()

		// 获取请求详情
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		// 记录请求详情
		// 看情况是否记录
		global.Logger.Info(fmt.Sprintf("[%s]", constant.ProjectName),
			zap.String("time", end.Format(time.DateTime)),
			zap.Int("status", statusCode),
			zap.Duration("latency", latency),
			zap.String("clientIP", clientIP),
			zap.String("method", method),
			zap.String("path", path),
		)
	}
}
