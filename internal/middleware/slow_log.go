package middleware

import (
	"gin-scaffold/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func SlowLogMiddleware(threshold time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求
		c.Next()

		duration := time.Since(start)

		// 如果超过阈值就记录慢日志
		if duration > threshold {
			logger.SlowLogger.Warn("慢请求",
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.Duration("duration", duration),
				zap.Int("status", c.Writer.Status()),
				zap.String("client_ip", c.ClientIP()),
			)
		}
	}
}
