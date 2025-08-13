package system

import (
	"gin-scaffold/internal/app/types/common"
	"time"
)

// DepartmentCreateReq 创建部门请求
type DepartmentCreateReq struct {
	DeptName string `json:"deptName" binding:"required"` // 部门名称
	ParentId *int64 `json:"parentId"`                    // 上级部门Id
	Status   int8   `json:"status" binding:"required"`   // 状态：1启用 0禁用
}

// DepartmentCreateResp 创建部门响应
type DepartmentCreateResp struct {
	Id int64 `json:"id"` // 部门Id
}

// DepartmentUpdateReq 更新部门请求
type DepartmentUpdateReq struct {
	Id       int64  `json:"id" binding:"required"`       // 部门Id
	DeptName string `json:"deptName" binding:"required"` // 部门名称
	ParentId int64  `json:"parentId"`                    // 上级部门Id
	Status   int8   `json:"status" binding:"required"`   // 状态：1启用 0禁用
}

// DepartmentUpdateResp 更新部门响应
type DepartmentUpdateResp struct{}

// DepartmentDeleteResp 删除部门响应
type DepartmentDeleteResp struct{}

// DepartmentDataResp 部门数据响应
type DepartmentDataResp struct {
	Id       int64  `json:"id"`       // 部门Id
	DeptName string `json:"deptName"` // 部门名称
	//TenantId  int64      `json:"tenantId"`  // 租户Id
	ParentId  *int64     `json:"parentId"`  // 上级部门Id
	Status    int8       `json:"status"`    // 状态：1启用 0禁用
	CreatedAt time.Time  `json:"createdAt"` // 创建时间
	CreatedBy int64      `json:"createdBy"` // 创建人Id
	UpdatedAt time.Time  `json:"updatedAt"` // 更新时间
	UpdatedBy int64      `json:"updatedBy"` // 更新人Id
	DeletedAt *time.Time `json:"deletedAt"` // 删除时间
	DeletedBy int64      `json:"deletedBy"` // 删除人Id
}

// DepartmentDataListResp 部门列表响应
type DepartmentDataListResp struct {
	common.ListResp
	Total int64                `json:"total"` // 总数
	List  []DepartmentDataResp `json:"list"`  // 列表
}

// DepartmentQueryReq 部门查询请求
type DepartmentQueryReq struct {
	common.ListReq
	DeptName string `form:"deptName"` // 部门名称
	Status   *int8  `form:"status"`   // 状态
}

// DepartmentTreeResp 部门树响应
type DepartmentTreeResp struct {
	Id       int64  `json:"id"`       // 部门Id
	DeptName string `json:"deptName"` // 部门名称
	//TenantId int64                `json:"tenantId"`           // 租户Id
	ParentId int64                `json:"parentId"`           // 上级部门Id
	Status   int8                 `json:"status"`             // 状态：1启用 0禁用
	Children []DepartmentTreeResp `json:"children,omitempty"` // 子部门
}
