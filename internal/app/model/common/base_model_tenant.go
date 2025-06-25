package common

import (
	"context"
	"gin-scaffold/internal/middleware/metadata"
	"gorm.io/gorm"
)

type BaseModelOnlyTenant struct {
	Id       int64 `json:"id"`
	TenantId int64 `json:"tenantId"`
}

// BeforeCreate 创建前钩子
func (m *BaseModelOnlyTenant) BeforeCreate(tx *gorm.DB) error {
	// 从上下文中获取用户Id
	if ctx, ok := tx.Statement.Context.(context.Context); ok {
		//if tenantId, err := metadata.GetTenantId(ctx); err == nil {
		m.TenantId = metadata.GetTenantId(ctx)
	}
	return nil
}
