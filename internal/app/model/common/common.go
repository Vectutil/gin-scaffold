package common

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int64          `gorm:"column:id" json:"id"`                // 主键
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64          `gorm:"column:created_by" json:"createdBy"` // 创建人ID
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
	UpdatedBy int64          `gorm:"column:updated_by" json:"updatedBy"` // 更新人ID
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
	DeletedBy int64          `gorm:"column:deleted_by" json:"deletedBy"` // 删除人ID
}
