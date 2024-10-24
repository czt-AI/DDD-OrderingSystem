package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware 错误恢复中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 捕获panic
		defer func() {
			if r := recover(); r != nil {
				// 记录错误
				log.Printf("Recovered from panic: %v", r)

				// 设置HTTP状态码为500
				c AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}()

		// 继续处理请求
		c.Next()
	}
}
