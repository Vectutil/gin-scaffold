package system

import "gin-scaffold/internal/app/model/common"

// User 用户表
type User struct {
	common.BaseModel
	Username    string `gorm:"column:username" json:"username"`         // 用户名
	Password    string `gorm:"column:password" json:"-"`                // 密码
	FullName    string `gorm:"column:full_name" json:"fullName"`        // 全名
	Avatar      string `gorm:"column:avatar" json:"avatar"`             // 头像URL
	Email       string `gorm:"column:email" json:"email"`               // 邮箱
	Phone       string `gorm:"column:phone" json:"phone"`               // 手机号
	DeptId      int64  `gorm:"column:dept_id" json:"deptId"`            // 所属主部门Id
	Status      int    `gorm:"column:status" json:"status"`             // 状态：1启用 0禁用
	LoginCount  int    `gorm:"column:login_count" json:"loginCount"`    // 登录次数
	LastLoginAt int64  `gorm:"column:last_login_at" json:"lastLoginAt"` // 最后登录时间
	LastLoginIP string `gorm:"column:last_login_ip" json:"lastLoginIp"` // 最后登录IP地址
	OrgId       int64  `gorm:"column:org_id" json:"orgId"`              // 组织Id
	Remark      string `gorm:"column:remark" json:"remark"`             // 备注信息
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}
