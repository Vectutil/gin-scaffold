package metadata

import (
	"time"
)

type IOperator interface {
	GetUserId() uint64                      // 获取用户Id
	GetEmpId() uint64                       // 获取用户员工id
	GetUserName() string                    // 获取用户名称
	GetDepartmentId() uint64                // 获取用户所属部门Id
	IsSaasCompany() bool                    // 所属公司是否为saas组织
	IsAvailable() bool                      // 用户是否可用
	GetToken() string                       // 获取用户登录token
	GetSubDepartmentId() []uint64           // 获取用户所在部门的子部门
	GetRoleAccessIds() []uint64             // 获取用户所属角色权限列表
	GetLoginTime() time.Time                // 获取登录时间
	GetMenuIds() (r []uint64)               // 获取菜单Id
	GetResourceRouterNames() (r []string)   // 获取前端路由名称
	GetButtonCodes() (r []string)           // 获取前端按钮编号
	GetMPResourceRouterNames() (r []string) // 获取前端内部商城路由名称
	GetMPButtonCodes() (r []string)         // 获取前端内部商城按钮编号
	GetWarehouseIds() (r []uint64)          // 获取仓库权限
	GetSaleSystemIds() (r []uint64)         // 获取营销体系权限
	GetBizUnitIds() (r []uint64)            // 获取往来单位权限
}

type IMallOperator interface {
	IOperator
}
