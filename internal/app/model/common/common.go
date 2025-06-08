package common

import (
	"context"
	"gin-scaffold/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int64          `gorm:"column:id" json:"id"` // 主键
	TenantID  int64          `json:"tenantId" gorm:"not null;default:0;comment:租户ID"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64          `gorm:"column:created_by" json:"createdBy"` // 创建人ID
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
	UpdatedBy int64          `gorm:"column:updated_by" json:"updatedBy"` // 更新人ID
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
	DeletedBy int64          `gorm:"column:deleted_by" json:"deletedBy"` // 删除人ID
}

// TenantScope 租户作用域
func TenantScope(ctx context.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if tenantID, err := utils.GetTenantIDFromContext(ctx); err == nil {
			return db.Where("tenant_id = ?", tenantID)
		}
		return db
	}
}

// NotDeletedScope 未删除作用域
func NotDeletedScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at IS NULL")
	}
}

// BeforeCreate 创建前钩子
func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	// 从上下文中获取租户ID和用户ID
	if ctx, ok := tx.Statement.Context.(context.Context); ok {
		if tenantID, err := utils.GetTenantIDFromContext(ctx); err == nil {
			m.TenantID = tenantID
		}
		if userID, err := utils.GetUserIDFromContext(ctx); err == nil {
			m.CreatedBy = userID
			m.UpdatedBy = userID
		}
	}
	return nil
}

// BeforeUpdate 更新前钩子
func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	// 从上下文中获取用户ID
	if ctx, ok := tx.Statement.Context.(context.Context); ok {
		if userID, err := utils.GetUserIDFromContext(ctx); err == nil {
			m.UpdatedBy = userID
		}
	}
	return nil
}

// BeforeDelete 删除前钩子
func (m *BaseModel) BeforeDelete(tx *gorm.DB) error {
	// 从上下文中获取用户ID
	if ctx, ok := tx.Statement.Context.(context.Context); ok {
		if userID, err := utils.GetUserIDFromContext(ctx); err == nil {
			m.DeletedBy = userID
		}
	}
	return nil
}
