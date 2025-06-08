package system

import (
	"gin-scaffold/internal/app/model/common"
	"time"
)

// Department 部门模型
type Department struct {
	common.BaseModel
	DeptName  string     `json:"deptName" gorm:"size:100;not null;comment:部门名称"`
	TenantID  int64      `json:"tenantId" gorm:"not null;default:0;comment:租户ID"`
	ParentID  *int64     `json:"parentId" gorm:"comment:上级部门ID，NULL 表示顶级"`
	Status    int8       `json:"status" gorm:"not null;default:1;comment:状态：1启用 0禁用"`
	CreatedAt time.Time  `json:"createdAt" gorm:"not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	CreatedBy int64      `json:"createdBy" gorm:"not null;default:0;comment:创建人ID"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"not null;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP;comment:更新时间"`
	UpdatedBy int64      `json:"updatedBy" gorm:"not null;default:0;comment:更新人ID"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"comment:删除时间"`
	DeletedBy int64      `json:"deletedBy" gorm:"not null;default:0;comment:删除人ID"`
}

// TableName 设置表名
func (Department) TableName() string {
	return "dept"
}

// DepartmentTree 部门树结构
type DepartmentTree struct {
	ID       int64            `json:"id"`
	DeptName string           `json:"deptName"`
	TenantID int64            `json:"tenantId"`
	ParentID *int64           `json:"parentId"`
	Status   int8             `json:"status"`
	Children []DepartmentTree `json:"children,omitempty"`
}
