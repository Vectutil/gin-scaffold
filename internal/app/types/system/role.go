package system

import (
	"gin-scaffold/internal/app/types/common"
	"time"
)

// RoleCreateReq 创建角色请求
type RoleCreateReq struct {
	Name        string `json:"name" binding:"required"` // 角色名称
	Code        string `json:"code" binding:"required"` // 角色编码
	Description string `json:"description"`             // 描述
	DataScope   int8   `json:"dataScope"`              // 数据范围
	Status      int8   `json:"status"`                 // 状态
}

// RoleCreateResp 创建角色响应
type RoleCreateResp struct {
	ID int64 `json:"id"` // 角色ID
}

// RoleUpdateReq 更新角色请求
type RoleUpdateReq struct {
	ID          int64  `json:"id" binding:"required"` // 角色ID
	Name        string `json:"name" binding:"required"` // 角色名称
	Code        string `json:"code" binding:"required"` // 角色编码
	Description string `json:"description"`             // 描述
	DataScope   int8   `json:"dataScope"`              // 数据范围
	Status      int8   `json:"status"`                 // 状态
}

// RoleUpdateResp 更新角色响应
type RoleUpdateResp struct{}

// RoleDeleteResp 删除角色响应
type RoleDeleteResp struct{}

// RoleDataResp 角色数据响应
type RoleDataResp struct {
	ID          int64      `json:"id"`          // 角色ID
	Name        string     `json:"name"`        // 角色名称
	Code        string     `json:"code"`        // 角色编码
	Description string     `json:"description"` // 描述
	DataScope   int8       `json:"dataScope"`   // 数据范围
	Status      int8       `json:"status"`      // 状态
	CreatedAt   time.Time  `json:"createdAt"`   // 创建时间
	CreatedBy   int64      `json:"createdBy"`   // 创建人
	UpdatedAt   time.Time  `json:"updatedAt"`   // 更新时间
	UpdatedBy   int64      `json:"updatedBy"`   // 更新人
	DeletedAt   *time.Time `json:"deletedAt"`   // 删除时间
	DeletedBy   int64      `json:"deletedBy"`   // 删除人
}

// RoleDataListResp 角色列表响应
type RoleDataListResp struct {
	common.ListResp
	Total int64           `json:"total"` // 总数
	List  []RoleDataResp  `json:"list"`  // 列表
}

// RoleQueryReq 角色查询请求
type RoleQueryReq struct {
	common.ListReq
	Name   string `form:"name"`   // 角色名称
	Code   string `form:"code"`   // 角色编码
	Status *int8  `form:"status"` // 状态
} 