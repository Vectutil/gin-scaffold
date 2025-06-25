package system

import (
	"gin-scaffold/internal/app/model/common"
)

// UserRoleRel 用户角色关系模型
type UserRoleRel struct {
	common.BaseModelOnlyTenant
	UserId int64 `json:"userId"` // 用户Id
	RoleId int64 `json:"roleId"` // 角色Id
}

// TableName 设置表名
func (UserRoleRel) TableName() string {
	return "user_role_rel"
}
