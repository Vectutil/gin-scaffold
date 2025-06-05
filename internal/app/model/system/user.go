package system

import "time"

// User 用户表
type User struct {
	ID          int64     `gorm:"column:id" json:"id"`                     // 主键
	Username    string    `gorm:"column:username" json:"username"`         // 用户名
	Password    string    `gorm:"column:password" json:"-"`                // 密码
	FullName    string    `gorm:"column:full_name" json:"fullName"`        // 全名
	Avatar      string    `gorm:"column:avatar" json:"avatar"`             // 头像URL
	Email       string    `gorm:"column:email" json:"email"`               // 邮箱
	Phone       string    `gorm:"column:phone" json:"phone"`               // 手机号
	DeptID      int64     `gorm:"column:dept_id" json:"deptId"`            // 所属主部门ID
	Status      int       `gorm:"column:status" json:"status"`             // 状态：1启用 0禁用
	LoginCount  int       `gorm:"column:login_count" json:"loginCount"`    // 登录次数
	LastLoginAt int64     `gorm:"column:last_login_at" json:"lastLoginAt"` // 最后登录时间
	LastLoginIP string    `gorm:"column:last_login_ip" json:"lastLoginIp"` // 最后登录IP地址
	TenantID    int64     `gorm:"column:tenant_id" json:"tenantId"`        // 租户ID
	OrgID       int64     `gorm:"column:org_id" json:"orgId"`              // 组织ID
	Remark      string    `gorm:"column:remark" json:"remark"`             // 备注信息
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`      // 创建时间
	CreatedBy   int64     `gorm:"column:created_by" json:"createdBy"`      // 创建人ID
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`      // 更新时间
	UpdatedBy   int64     `gorm:"column:updated_by" json:"updatedBy"`      // 更新人ID
	DeletedAt   time.Time `gorm:"column:deleted_at" json:"deletedAt"`      // 删除时间
	DeletedBy   int64     `gorm:"column:deleted_by" json:"deletedBy"`      // 删除人ID
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}
