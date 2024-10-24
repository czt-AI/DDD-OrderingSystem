package middleware

import (
	"github.com/gin-gonic/gin"
)

// ResponseMiddleware 响应中间件
func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 设置默认响应格式
		c.Writer.Header().Set("Content-Type", "application/json")

		// 获取响应数据
		var response gin.H
		if err := c.ShouldBindJSON(&response); err != nil {
			response = gin.H{"error": "Invalid request payload"}
		}

		// 设置响应状态码
		if c.Writer.Status() == 0 {
			c.Writer.WriteHeader(http.StatusOK)
		}

		// 发送响应
		c.JSON(c.Writer.Status(), response)
	}
}
