package utils

import (
	"context"
	"errors"
	systype "gin-scaffold/internal/app/types/system"
	"github.com/gin-gonic/gin"
)

const (
	UserIDKey   = "userID"
	TenantIDKey = "tenantID"
	DeptIdKey   = "deptId"
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
func GetUserIDFromContext(ctx context.Context) (int64, error) {
	if ctx == nil {
		return 0, errors.New("context is nil")
	}
	if userID, ok := ctx.Value(UserIDKey).(int64); ok {
		return userID, nil
	}
	return 0, errors.New("userID not found in context")
}

// GetTenantIDFromContext 从上下文中获取租户ID
func GetTenantIDFromContext(ctx context.Context) (int64, error) {
	if ctx == nil {
		return 0, errors.New("context is nil")
	}
	if tenantID, ok := ctx.Value(TenantIDKey).(int64); ok {
		return tenantID, nil
	}
	return 0, errors.New("tenantID not found in context")
}

// WithUserID 将用户ID添加到上下文
func WithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// WithTenantID 将租户ID添加到上下文
func WithTenantID(ctx context.Context, tenantID int64) context.Context {
	return context.WithValue(ctx, TenantIDKey, tenantID)
}
