package system

import (
	"gin-scaffold/internal/app/model/common"
)

// Role 角色模型
type Role struct {
	common.BaseModel
	Name        string `json:"name" gorm:"size:50;not null;default:'';comment:角色名称"`
	Code        string `json:"code" gorm:"size:50;not null;default:'';comment:角色编码"`
	Description string `json:"description" gorm:"size:255;default:'';comment:描述"`
	DataScope   int8   `json:"dataScope" gorm:"default:4;comment:数据范围: 1-全部, 2-本部门, 3-本部门及子部门, 4-本人, 5-自定义"`
	Status      int8   `json:"status" gorm:"default:1;comment:状态 1启用 0禁用"`
}

// TableName 设置表名
func (Role) TableName() string {
	return "role"
}
