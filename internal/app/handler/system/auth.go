package system

import (
	"crypto/rand"
	"encoding/base64"
	syslogic "gin-scaffold/internal/app/logic/system"
	"gin-scaffold/internal/app/response"
	systype "gin-scaffold/internal/app/types/system"
	"gin-scaffold/pkg/mysql"
	"gin-scaffold/pkg/redis"
	"time"

	"github.com/gin-gonic/gin"
)

// AuthHandler 认证处理器
type AuthHandler struct{}

// NewAuthHandler 创建认证处理器实例
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// Login 用户登录
// @title 用户登录
// @Summary 用户登录接口
// @Description 用户登录并获取访问令牌
// @Tags 认证管理
// @Accept json
// @Produce json
// @Param request body LoginReq true "登录请求参数"
// @Success 200 {object} LoginResp "成功返回"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 401 {object} response.Response "认证失败"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var (
		err       error
		errCode   int
		req       systype.LoginReq
		res       = &systype.LoginResp{}
		userLogic = syslogic.NewUserLogic(mysql.GetDB())
	)

	defer func() {
		response.HandleDefault(c, res)(&err, &errCode)
	}()

	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}

	// TODO: 验证用户名和密码
	// 这里应该调用 userLogic 进行实际的用户验证
	userLogic
	// 为了示例，我们假设验证通过

	// 生成随机token
	token := generateToken()

	// 将token存储在Redis中，设置过期时间为24小时
	err = redis.GetClient().Set(c.Request.Context(), "token:"+token, "1", 24*time.Hour)
	if err != nil {
		return
	}

	res.Token = token
}

// generateToken 生成随机token
func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
