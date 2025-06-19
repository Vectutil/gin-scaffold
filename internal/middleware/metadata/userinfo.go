package metadata

import (
	"context"
	"gin-scaffold/pkg/utils"
)

const (
	UserName     = "user_name"
	UserID       = "user_id"
	CompanyID    = "company_id"
	DepartmentID = "department_id"
	BizUnitID    = "biz_unit_id"
	SaleSystemID = "sale_system_id"
	TenantID     = "tenant_id"

	LoginInfo          = "login_info"
	DataSeparate       = "data_separate"
	IsAllowUpdateOrder = "is_allow_update_order"
	IsAllowCancelOther = "is_allow_cancel_other"
	IsAllowAuditSelf   = "is_allow_audit_self"
)

func GetUserName(ctx context.Context) string {
	return GetMD(ctx, UserName)
}

func GetUserID(ctx context.Context) int64 {
	str := GetMD(ctx, UserID)
	if str == "" {
		return 0
	}
	return utils.String2int64(str)
}

func GetCompanyID(ctx context.Context) uint64 {
	str := GetMD(ctx, CompanyID)
	if str == "" {
		return 0
	}
	return utils.String2Uint64(str)
}

func GetDepartmentID(ctx context.Context) uint64 {
	str := GetMD(ctx, DepartmentID)
	if str == "" {
		return 0
	}
	return utils.String2Uint64(str)
}

func GetBizUnitID(ctx context.Context) uint64 {
	str := GetMD(ctx, BizUnitID)
	if str == "" {
		return 0
	}
	return utils.String2Uint64(str)
}

func GetSaleSystemID(ctx context.Context) uint64 {
	str := GetMD(ctx, SaleSystemID)
	if str == "" {
		return 0
	}
	return utils.String2Uint64(str)
}

func GetTenantID(ctx context.Context) int64 {
	str := GetMD(ctx, TenantID)
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
