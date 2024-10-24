package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求
		c.Next()

		// 计算处理时间
		latency := time.Since(start)
		statusCode := c.Writer.Status()

		// 打印日志
		fmt.Printf(
			"%s | %3d | %13v | %s | %s\n",
			c.Request.Method,
			statusCode,
			latency,
			c.Request.URL.Path,
			c.ClientIP(),
		)
	}
}
