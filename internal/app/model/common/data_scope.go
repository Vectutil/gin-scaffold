package common

import (
	"context"
	"gin-scaffold/internal/middleware/metadata"
	"gorm.io/gorm"
)

// TenantScope 租户作用域
func TenantScope(ctx context.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//if tenantID, err := utils.GetTenantIDFromContext(ctx); err == nil {
		if metadata.GetTenantID(ctx) != 0 {
			return db.Where("tenant_id = ?", metadata.GetTenantID(ctx))
		}
		//}
		return db
	}
}

func UserScope(ctx context.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//if userId, err := utils.GetUserIDFromContext(ctx); err == nil {
		//}
		if metadata.GetUserID(ctx) != 0 {
			return db.Where("created_by = ?", metadata.GetUserID(ctx))
		}
		return db
	}
}

func DeptScope() {

}
