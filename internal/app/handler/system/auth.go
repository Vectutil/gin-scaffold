package system

import (
	"encoding/json"
	syslogic "gin-scaffold/internal/app/logic/system"
	"gin-scaffold/internal/app/response"
	systype "gin-scaffold/internal/app/types/system"
	"gin-scaffold/pkg/mysql"
	"gin-scaffold/pkg/redis"
	"gin-scaffold/pkg/utils"
	"net/http"
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
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var (
		err       error
		req       systype.LoginReq
		res       = &systype.LoginResp{}
		userLogic = syslogic.NewUserLogic(mysql.GetDB())
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}

	// 这里应该调用 userLogic 进行实际的用户验证
	userInfo, err := userLogic.CheckForLogin(c, req.Phone, req.Password)
	if err != nil {
		return
	}

	// 生成随机token
	accToken, _ := utils.GenerateAccessToken(userInfo.Id, utils.RoleUser)
	refToken, _ := utils.GenerateRefreshToken(userInfo.Id)

	userData, _ := json.Marshal(userInfo)
	// 将token存储在Redis中，设置过期时间为7*24小时
	err = redis.GetClient().Set(c.Request.Context(), "accToken:"+accToken, userData, 2*time.Hour)
	if err != nil {
		return
	}
	err = redis.GetClient().Set(c.Request.Context(), "refToken:"+refToken, userData, 7*24*time.Hour)
	if err != nil {
		return
	}

	res.AccessToken = accToken
	res.RefreshToken = refToken

}

// RefreshToken 刷新访问令牌
// @title 刷新访问令牌
// @Summary 使用刷新令牌获取新的访问令牌
// @Description 使用刷新令牌获取新的访问令牌
// @Tags 认证管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 刷新令牌"
// @Success 200 {object} map[string]string "成功返回"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 401 {object} response.Response "认证失败"
// @Router /refresh [post]
func RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	claims, err := utils.ParseToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

	// 生成新的 access token
	newAccessToken, _ := utils.GenerateAccessToken(claims.UserId, claims.Role)

	// 在响应头中添加 Authorization
	c.Header("Authorization", "Bearer "+newAccessToken)

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}

// Register 用户注册
// @title 用户注册
// @Summary 用户注册接口
// @Description 新用户注册并获取访问令牌
// @Tags 认证管理
// @Accept json
// @Produce json
// @Param request body systype.RegisterReq true "注册请求参数"
// @Success 200 {object} systype.RegisterResp "成功返回"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var (
		err       error
		req       systype.RegisterReq
		res       = &systype.RegisterResp{}
		userLogic = syslogic.NewUserLogic(mysql.GetDB())
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}

	// 验证两次密码是否一致
	if req.Password != req.ConfirmPassword {
		err = response.NewError(http.StatusBadRequest, "两次输入的密码不一致")
		return
	}

	// 创建用户
	createReq := &systype.UserCreateReq{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Email:    req.Email,
		FullName: req.FullName,
		Status:   1, // 默认启用状态
	}

	if err = userLogic.CreateForRegister(c.Request.Context(), createReq); err != nil {
		return
	}
}
