package middleware

import (
	"encoding/json"
	"gin-scaffold/pkg/redis"
	"gin-scaffold/pkg/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

		var userInfo MetaData
		_ = json.Unmarshal([]byte(data), &userInfo)
		if userInfo.ID != claims.UserID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token mismatch"})
			return
		}
		// 3. 设置到 Gin 上下文中
		c.Set("user", userInfo)
		c.Next()
	}
}

type MetaData struct {
	ID          int64     `json:"id"`          // 主键
	Username    string    `json:"username"`    // 用户名
	Password    string    `json:"-"`           // 密码
	FullName    string    `json:"fullName"`    // 全名
	Avatar      string    `json:"avatar"`      // 头像URL
	Email       string    `json:"email"`       // 邮箱
	Phone       string    `json:"phone"`       // 手机号
	DeptID      int64     `json:"deptId"`      // 所属主部门ID
	Status      int       `json:"status"`      // 状态：1启用 0禁用
	LoginCount  int       `json:"loginCount"`  // 登录次数
	LastLoginAt int64     `json:"lastLoginAt"` // 最后登录时间
	LastLoginIP string    `json:"lastLoginIp"` // 最后登录IP地址
	TenantID    int64     `json:"tenantId"`    // 租户ID
	OrgID       int64     `json:"orgId"`       // 组织ID
	Remark      string    `json:"remark"`      // 备注信息
	CreatedAt   time.Time `json:"createdAt"`   // 创建时间
	CreatedBy   int64     `json:"createdBy"`   // 创建人ID
	UpdatedAt   time.Time `json:"updatedAt"`   // 更新时间
	UpdatedBy   int64     `json:"updatedBy"`   // 更新人ID
}
