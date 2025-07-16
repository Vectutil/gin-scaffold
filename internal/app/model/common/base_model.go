package common

import (
	"context"
	"gin-scaffold/internal/middleware/metadata"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id int64 `gorm:"column:id" json:"id"` // 主键
	//TenantId  int64          `json:"tenantId" gorm:"not null;default:0;comment:租户Id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64          `gorm:"column:created_by" json:"createdBy"` // 创建人Id
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
	UpdatedBy int64          `gorm:"column:updated_by" json:"updatedBy"` // 更新人Id
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
	DeletedBy int64          `gorm:"column:deleted_by" json:"deletedBy"` // 删除人Id
}

// BeforeCreate 创建前钩子
func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	// 从上下文中获取租户Id和用户Id
	if ctx, ok := tx.Statement.Context.(context.Context); ok {
		//if tenantId, err := utils.GetTenantIdFromContext(ctx); err == nil {
		//m.TenantId = metadata.GetTenantId(ctx)
		//}
		//if userId, err := utils.GetUserIdFromContext(ctx); err == nil {
		m.CreatedBy = metadata.GetUserId(ctx)
		m.UpdatedBy = metadata.GetUserId(ctx)
		//}
	}
	return nil
}

// BeforeUpdate 更新前钩子
func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	// 从上下文中获取用户Id
	if ctx, ok := tx.Statement.Context.(context.Context); ok {
		//if userId, err := utils.GetUserIdFromContext(ctx); err == nil {
		m.UpdatedBy = metadata.GetUserId(ctx)
		//}
	}
	return nil
}

// BeforeDelete 删除前钩子
func (m *BaseModel) BeforeDelete(tx *gorm.DB) error {
	// 从上下文中获取用户Id
	if ctx, ok := tx.Statement.Context.(context.Context); ok {
		//if userId, err := utils.GetUserIdFromContext(ctx); err == nil {
		m.DeletedBy = metadata.GetUserId(ctx)
		//}
	}
	return nil
}
