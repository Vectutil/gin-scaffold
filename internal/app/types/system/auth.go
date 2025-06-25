package system

// LoginReq 登录请求参数
type LoginReq struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResp 登录响应
type LoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthRedisInfo struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	DeptId   int64  `json:"dept_id"`
	TenantId int64  `json:"tenant_id"`
	OrgId    int64  `json:"org_id"`
	Status   int    `json:"status"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

// RegisterReq 注册请求
type RegisterReq struct {
	Username        string `json:"username" binding:"required"`         // 用户名
	Password        string `json:"password" binding:"required"`         // 密码
	ConfirmPassword string `json:"confirm_password" binding:"required"` // 确认密码
	Phone           string `json:"phone" binding:"required"`            // 手机号
	Email           string `json:"email"`                               // 邮箱
	FullName        string `json:"full_name"`                           // 姓名
}

// RegisterResp 注册响应
type RegisterResp struct {
	Id          int64  `json:"id"`           // 用户Id
	Username    string `json:"username"`     // 用户名
	Phone       string `json:"phone"`        // 手机号
	Email       string `json:"email"`        // 邮箱
	FullName    string `json:"full_name"`    // 姓名
	AccessToken string `json:"access_token"` // 访问令牌
}
