package middleware

import (
	"gin-scaffold/pkg/redis"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 是一个 Gin 中间件，用于验证用户身份
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// 检查 Bearer token 格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		token := parts[1]

		// 从 Redis 中验证 token
		exists, err := redis.GetClient().Exists(c.Request.Context(), "token:"+token).Result()
		if err != nil || exists == 0 {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 获取用户ID
		userID, err := redis.GetClient().Get(c.Request.Context(), "token:"+token).Result()
		if err != nil {
			c.JSON(401, gin.H{"error": "Failed to get user information"})
			c.Abort()
			return
		}

		// 将用户ID存储在上下文中
		c.Set("userID", userID)
		c.Next()
	}
} 