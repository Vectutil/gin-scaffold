package system

import "gin-scaffold/internal/app/types/common"

// UserRoleRelCreateReq 创建用户角色关系请求
type UserRoleRelCreateReq struct {
	UserID int64 `json:"userId" binding:"required"` // 用户ID
	RoleID int64 `json:"roleId" binding:"required"` // 角色ID
}

// UserRoleRelDeleteReq 删除用户角色关系请求
type UserRoleRelDeleteReq struct {
	UserID int64 `json:"userId" binding:"required"` // 用户ID
	RoleID int64 `json:"roleId" binding:"required"` // 角色ID
}

// UserRoleRelQueryReq 查询用户角色关系请求
type UserRoleRelQueryReq struct {
	UserID int64 `form:"userId"` // 用户ID
	RoleID int64 `form:"roleId"` // 角色ID
}

// UserRoleRelDataResp 用户角色关系数据响应
type UserRoleRelDataResp struct {
	ID       int64 `json:"id"`       // 主键
	TenantID int64 `json:"tenantId"` // 租户ID
	UserID   int64 `json:"userId"`   // 用户ID
	RoleID   int64 `json:"roleId"`   // 角色ID
}

// UserRoleRelDataListResp 用户角色关系列表响应
type UserRoleRelDataListResp struct {
	common.ListResp
	Total int64                `json:"total"` // 总数
	List  []UserRoleRelDataResp `json:"list"`  // 列表
} 