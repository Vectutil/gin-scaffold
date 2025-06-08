package utils

import (
	"gin-scaffold/internal/app/types/system"
	"github.com/gin-gonic/gin"
)

// GetUserFromContext 从上下文中获取用户信息
func GetUserFromContext(c *gin.Context) (*systype.UserDataResp, error) {
	user, exists := c.Get("user")
	if !exists {
		return nil, ErrUserNotFound
	}

	userInfo, ok := user.(*systype.UserDataResp)
	if !ok {
		return nil, ErrInvalidUserInfo
	}

	return userInfo, nil
}

// GetUserIDFromContext 从上下文中获取用户ID
func GetUserIDFromContext(c *gin.Context) (int64, error) {
	userInfo, err := GetUserFromContext(c)
	if err != nil {
		return 0, err
	}
	return userInfo.ID, nil
} 