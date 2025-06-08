package system

import "gin-scaffold/internal/app/model/common"

// UserRoleRel 用户角色关系模型
type UserRoleRel struct {
	common.BaseModel
	TenantID int64 `json:"tenantId" gorm:"not null;default:0;comment:租户ID"`
	UserID   int64 `json:"userId" gorm:"not null;comment:用户ID"`
	RoleID   int64 `json:"roleId" gorm:"not null;comment:角色ID"`
}

// TableName 设置表名
func (UserRoleRel) TableName() string {
	return "user_role_rel"
} 