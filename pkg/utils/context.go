package utils

import (
	"context"
	"errors"
	systype "gin-scaffold/internal/app/types/system"
	"github.com/gin-gonic/gin"
)

const (
	UserIdKey   = "userId"
	TenantIdKey = "tenantId"
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

// GetUserIdFromContext 从上下文中获取用户Id
func GetUserIdFromContext(ctx context.Context) (int64, error) {
	if ctx == nil {
		return 0, errors.New("context is nil")
	}
	if userId, ok := ctx.Value(UserIdKey).(int64); ok {
		return userId, nil
	}
	return 0, errors.New("userId not found in context")
}

// GetTenantIdFromContext 从上下文中获取租户Id
//func GetTenantIdFromContext(ctx context.Context) (int64, error) {
//	if ctx == nil {
//		return 0, errors.New("context is nil")
//	}
//	if tenantId, ok := ctx.Value(TenantIdKey).(int64); ok {
//		return tenantId, nil
//	}
//	return 0, errors.New("tenantId not found in context")
//}

// WithUserId 将用户Id添加到上下文
func WithUserId(ctx context.Context, userId int64) context.Context {
	return context.WithValue(ctx, UserIdKey, userId)
}

// WithTenantId 将租户Id添加到上下文
func WithTenantId(ctx context.Context, tenantId int64) context.Context {
	return context.WithValue(ctx, TenantIdKey, tenantId)
}
