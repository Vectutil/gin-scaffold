package middleware

import (
	"encoding/json"
	"gin-scaffold/internal/middleware/metadata"
	"gin-scaffold/pkg/redis"
	"gin-scaffold/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware 是一个 Gin 中间件，用于验证用户身份
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// 1. 校验 JWT
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// 2. 从 Redis 拿用户信息
		data, err := redis.GetClient().Get(c.Request.Context(), "accToken:"+token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired or not found"})
			return
		}

		var userInfo metadata.MetaData
		_ = json.Unmarshal([]byte(data), &userInfo)
		if userInfo.Id != claims.UserId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token mismatch"})
			return
		}
		// 3. 设置到 Gin 上下文中
		c.Set("user", userInfo)

		metadata.SetMetadataForUserInfo(c)

		c.Next()
	}
}
