package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiter sync.Map

// RateLimiterMiddleware 请求速率限制中间件
func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()
		limiter, ok := limiter.LoadOrStore(key, rate.NewLimiter(1, 5))
		if !ok {
			limiter = rate.NewLimiter(1, 5)
			limiter.(*rate.Limiter).SetLimit(1)
			limiter.(*rate.Limiter).SetBurst(5)
			limiter.Store(key, limiter)
		}

		if err := limiter.(*rate.Limiter).Wait(c); err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			return
		}

		c.Next()
	}
}
