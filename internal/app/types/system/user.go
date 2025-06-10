package system

import (
	"fmt"
	"gin-scaffold/internal/app/types/common"
	"strings"
	"time"
)

// User 用户信息
type User struct {
	ID          int64     `json:"id"`          // 主键
	Username    string    `json:"username"`    // 用户名
	Password    string    `json:"-"`           // 密码
	FullName    string    `json:"fullName"`    // 全名
	Avatar      string    `json:"avatar"`      // 头像URL
	Email       string    `json:"email"`       // 邮箱
	Phone       string    `json:"phone"`       // 手机号
	DeptID      int64     `json:"deptId"`      // 所属主部门ID
	Status      int       `json:"status"`      // 状态：1启用 0禁用
	LoginCount  int       `json:"loginCount"`  // 登录次数
	LastLoginAt int64     `json:"lastLoginAt"` // 最后登录时间
	LastLoginIP string    `json:"lastLoginIp"` // 最后登录IP地址
	TenantID    int64     `json:"tenantId"`    // 租户ID
	OrgID       int64     `json:"orgId"`       // 组织ID
	Remark      string    `json:"remark"`      // 备注信息
	CreatedAt   time.Time `json:"createdAt"`   // 创建时间
	CreatedBy   int64     `json:"createdBy"`   // 创建人ID
	UpdatedAt   time.Time `json:"updatedAt"`   // 更新时间
	UpdatedBy   int64     `json:"updatedBy"`   // 更新人ID
	DeletedAt   time.Time `json:"deletedAt"`   // 删除时间
	DeletedBy   int64     `json:"deletedBy"`   // 删除人ID
}

// UserCreateReq 创建用户请求
type UserCreateReq struct {
	common.BaseParam
	Username string  `json:"username" binding:"required"` // 用户名
	Password string  `json:"password" binding:"required"` // 密码
	FullName string  `json:"fullName"`                    // 全名
	Email    string  `json:"email"`                       // 邮箱
	Phone    string  `json:"phone"`                       // 手机号
	DeptID   int64   `json:"deptId"`                      // 所属主部门ID
	Status   int     `json:"status"`                      // 状态
	Remark   string  `json:"remark"`                      // 备注信息
	RoleIds  []int64 `json:"roleIds"`                     // 角色ID
}

func (u *UserCreateReq) Adjust() {
	if strings.TrimSpace(u.Username) == "" {
		u.Username = fmt.Sprintf("游客%d", time.Now().Unix())
	}
}

type UserCreateResp struct {
	ID int64 `json:"id"` // 主键
}

// UserUpdateReq 更新用户请求
type UserUpdateReq struct {
	ID       int64   `json:"id" binding:"required"`           // 主键
	Username string  `gorm:"column:username" json:"username"` // 用户名
	FullName string  `json:"fullName"`                        // 全名
	Email    string  `json:"email"`                           // 邮箱
	Phone    string  `json:"phone"`                           // 手机号
	DeptID   int64   `json:"deptId"`                          // 所属主部门ID
	Status   int     `json:"status"`                          // 状态
	Remark   string  `json:"remark"`                          // 备注信息
	RoleIds  []int64 `json:"roleIds"`                         // 角色ID
}
type UserUpdateResp struct {
	ID int64 `json:"id"` // 主键
}
type UserDeleteResp struct {
	ID int64 `json:"id"` // 主键
}

// UserQueryReq 查询用户请求
type UserQueryReq struct {
	Username string `form:"username"` // 用户名
	Email    string `form:"email"`    // 邮箱
	Phone    string `form:"phone"`    // 手机号
	DeptID   int64  `form:"deptId"`   // 所属主部门ID
	Status   int    `form:"status"`   // 状态
	common.ListReq
}

type UserDataResp struct {
	ID          int64     `json:"id"`          // 主键
	Username    string    `json:"username"`    // 用户名
	Password    string    `json:"-"`           // 密码
	FullName    string    `json:"fullName"`    // 全名
	Avatar      string    `json:"avatar"`      // 头像URL
	Email       string    `json:"email"`       // 邮箱
	Phone       string    `json:"phone"`       // 手机号
	DeptID      int64     `json:"deptId"`      // 所属主部门ID
	Status      int       `json:"status"`      // 状态：1启用 0禁用
	LoginCount  int       `json:"loginCount"`  // 登录次数
	LastLoginAt int64     `json:"lastLoginAt"` // 最后登录时间
	LastLoginIP string    `json:"lastLoginIp"` // 最后登录IP地址
	TenantID    int64     `json:"tenantId"`    // 租户ID
	OrgID       int64     `json:"orgId"`       // 组织ID
	Remark      string    `json:"remark"`      // 备注信息
	CreatedAt   time.Time `json:"createdAt"`   // 创建时间
	CreatedBy   int64     `json:"createdBy"`   // 创建人ID
	UpdatedAt   time.Time `json:"updatedAt"`   // 更新时间
	UpdatedBy   int64     `json:"updatedBy"`   // 更新人ID
}

type UserDataListResp struct {
	common.ListResp
	List []*UserDataResp `json:"list"`
}
