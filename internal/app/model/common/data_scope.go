package common

import (
	"context"
	"gin-scaffold/internal/middleware/metadata"
	"gorm.io/gorm"
)

// TenantScope 租户作用域
func TenantScope(ctx context.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//if tenantId, err := utils.GetTenantIdFromContext(ctx); err == nil {
		if metadata.GetTenantId(ctx) != 0 {
			return db.Where("tenant_id = ?", metadata.GetTenantId(ctx))
		}
		//}
		return db
	}
}

func UserScope(ctx context.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//if userId, err := utils.GetUserIdFromContext(ctx); err == nil {
		//}
		if metadata.GetUserId(ctx) != 0 {
			return db.Where("created_by = ?", metadata.GetUserId(ctx))
		}
		return db
	}
}

func DeptScope() {

}
