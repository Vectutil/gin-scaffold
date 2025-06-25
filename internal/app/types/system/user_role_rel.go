package system

import "gin-scaffold/internal/app/types/common"

// UserRoleRelCreateReq 创建用户角色关系请求
type UserRoleRelCreateReq struct {
	UserId int64 `json:"userId" binding:"required"` // 用户Id
	RoleId int64 `json:"roleId" binding:"required"` // 角色Id
}

// UserRoleRelDeleteReq 删除用户角色关系请求
type UserRoleRelDeleteReq struct {
	UserId int64 `json:"userId" binding:"required"` // 用户Id
	RoleId int64 `json:"roleId" binding:"required"` // 角色Id
}

// UserRoleRelQueryReq 查询用户角色关系请求
type UserRoleRelQueryReq struct {
	UserId int64 `form:"userId"` // 用户Id
	RoleId int64 `form:"roleId"` // 角色Id
}

// UserRoleRelDataResp 用户角色关系数据响应
type UserRoleRelDataResp struct {
	Id       int64 `json:"id"`       // 主键
	TenantId int64 `json:"tenantId"` // 租户Id
	UserId   int64 `json:"userId"`   // 用户Id
	RoleId   int64 `json:"roleId"`   // 角色Id
}

// UserRoleRelDataListResp 用户角色关系列表响应
type UserRoleRelDataListResp struct {
	common.ListResp
	Total int64                 `json:"total"` // 总数
	List  []UserRoleRelDataResp `json:"list"`  // 列表
}
