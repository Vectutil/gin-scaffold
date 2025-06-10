package system

import (
	"gin-scaffold/internal/app/model/common"
)

// UserRoleRel 用户角色关系模型
type UserRoleRel struct {
	common.BaseModel
	TenantID int64 `json:"tenantId"`
	UserID   int64 `json:"userId"`
	RoleID   int64 `json:"roleId"`
}

// TableName 设置表名
func (UserRoleRel) TableName() string {
	return "user_role_rel"
}
