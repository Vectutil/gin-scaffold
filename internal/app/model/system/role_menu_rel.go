package system

import (
	"gin-scaffold/internal/app/model/common"
)

// RoleMenuRel 角色权限映射表模型
type RoleMenuRel struct {
	common.BaseModelOnlyTenant
	RoleId    int64 `json:"roleId"`    // 角色Id
	MenuId    int64 `json:"menuId"`    // 菜单Id（按钮/菜单）
	ScopeType int8  `json:"scopeType"` // 数据范围: 1-全部, 2-本部门, 3-本部门及子部门, 4-本人
}

// TableName 设置表名
func (RoleMenuRel) TableName() string {
	return "role_menu_rel"
}
