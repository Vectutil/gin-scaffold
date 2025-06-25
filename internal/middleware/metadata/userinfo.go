package metadata

import (
	"context"
	"gin-scaffold/pkg/utils"
)

const (
	UserName     = "user_name"
	UserId       = "user_id"
	CompanyId    = "company_id"
	DepartmentId = "department_id"
	BizUnitId    = "biz_unit_id"
	SaleSystemId = "sale_system_id"
	TenantId     = "tenant_id"

	LoginInfo          = "login_info"
	DataSeparate       = "data_separate"
	IsAllowUpdateOrder = "is_allow_update_order"
	IsAllowCancelOther = "is_allow_cancel_other"
	IsAllowAuditSelf   = "is_allow_audit_self"
)

func GetUserName(ctx context.Context) string {
	return GetMD(ctx, UserName)
}

func GetUserId(ctx context.Context) int64 {
	str := GetMD(ctx, UserId)
	if str == "" {
		return 0
	}
	return utils.String2int64(str)
}

func GetCompanyId(ctx context.Context) uint64 {
	str := GetMD(ctx, CompanyId)
	if str == "" {
		return 0
	}
	return utils.String2Uint64(str)
}

func GetDepartmentId(ctx context.Context) uint64 {
	str := GetMD(ctx, DepartmentId)
	if str == "" {
		return 0
	}
	return utils.String2Uint64(str)
}

func GetBizUnitId(ctx context.Context) uint64 {
	str := GetMD(ctx, BizUnitId)
	if str == "" {
		return 0
	}
	return utils.String2Uint64(str)
}

func GetSaleSystemId(ctx context.Context) uint64 {
	str := GetMD(ctx, SaleSystemId)
	if str == "" {
		return 0
	}
	return utils.String2Uint64(str)
}

func GetTenantId(ctx context.Context) int64 {
	str := GetMD(ctx, TenantId)
	if str == "" {
		return 0
	}
	return utils.String2int64(str)
}

func GetLoginInfo(ctx context.Context) IOperator {
	operator, ok := ctx.Value(LoginInfo).(IOperator)
	if ok {
		return operator
	}

	return nil
}

func GetIsAllowUpdateOrder(ctx context.Context) bool {
	isAllowUpdateOrder := GetMD(ctx, IsAllowUpdateOrder)
	if isAllowUpdateOrder == "1" {
		return true
	} else {
		return false
	}
}

func GetIsAllowCancelOther(ctx context.Context) bool {
	isAllowCancelOther := GetMD(ctx, IsAllowCancelOther)
	if isAllowCancelOther == "1" {
		return true
	} else {
		return false
	}
}

func GetIsAllowAuditSelf(ctx context.Context) bool {
	isAllowAuditSelf := GetMD(ctx, IsAllowAuditSelf)
	if isAllowAuditSelf == "1" {
		return true
	} else {
		return false
	}
}
