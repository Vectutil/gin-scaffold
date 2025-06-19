package common

import (
	"context"
	"gin-scaffold/internal/middleware/metadata"
	"gorm.io/gorm"
)

type BaseModelOnlyTenant struct {
	TenantID int64 `json:"tenantId" gorm:"not null;default:0;comment:租户ID"`
}

// BeforeCreate 创建前钩子
func (m *BaseModelOnlyTenant) BeforeCreate(tx *gorm.DB) error {
	// 从上下文中获取用户ID
	if ctx, ok := tx.Statement.Context.(context.Context); ok {
		//if tenantID, err := metadata.GetTenantID(ctx); err == nil {
		m.TenantID = metadata.GetTenantID(ctx)
	}
	return nil
}
