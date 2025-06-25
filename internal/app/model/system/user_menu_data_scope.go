package system

import (
	"gin-scaffold/internal/app/model/common"
)

// UserMenuDataScope 用户在某功能下可访问的部门（个性化权限）模型
type UserMenuDataScope struct {
	common.BaseModelOnlyTenant
	UserId  int64  `json:"userId"`  // 用户Id
	MenuId  int64  `json:"menuId"`  // 菜单Id
	DeptIds string `json:"deptIds"` // 部门Ids
}

// TableName 设置表名
func (UserMenuDataScope) TableName() string {
	return "user_menu_data_scope"
}
