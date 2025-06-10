package system

import (
	"gin-scaffold/internal/app/model/common"
)

// Department 部门模型
type Department struct {
	common.BaseModel
	Name     string `json:"name"`      // 部门名称
	ParentID int64  `json:"parent_id"` // 上级部门ID，NULL 表示顶级
	Level    int8   `json:"level"`     // 深度
	Status   int8   `json:"status"`    // 状态：1启用 0禁用
}

// TableName 设置表名
func (Department) TableName() string {
	return "department"
}

type DepartmentTree struct {
	common.BaseModel
	Name     string           `json:"name"`      // 部门名称
	ParentID int64            `json:"parent_id"` // 上级部门ID，NULL 表示顶级
	Level    int8             `json:"level"`     // 深度
	Status   int8             `json:"status"`    // 状态：1启用 0禁用
	Children []DepartmentTree `json:"children,omitempty"`
}
