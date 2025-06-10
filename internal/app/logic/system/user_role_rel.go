package system

import (
	"context"
	sysdao "gin-scaffold/internal/app/dao/system"
	sysmodel "gin-scaffold/internal/app/model/system"
	systype "gin-scaffold/internal/app/types/system"
	"gorm.io/gorm"
)

// UserRoleRelLogic 用户角色关系逻辑
type UserRoleRelLogic struct {
	dao *sysdao.UserRoleRelDao
}

// NewUserRoleRelLogic 创建用户角色关系逻辑实例
func NewUserRoleRelLogic(db *gorm.DB) *UserRoleRelLogic {
	return &UserRoleRelLogic{dao: sysdao.NewUserRoleRelDao(db)}
}

// Create 创建用户角色关系
func (l *UserRoleRelLogic) Create(ctx context.Context, req *systype.UserRoleRelCreateReq, tenantID int64) error {
	rel := &sysmodel.UserRoleRel{
		TenantID: tenantID,
		UserID:   req.UserID,
		RoleID:   req.RoleID,
	}
	return l.dao.Create(ctx, rel)
}

// Delete 删除用户角色关系
func (l *UserRoleRelLogic) Delete(ctx context.Context, userID, roleID int64) error {
	return l.dao.Delete(ctx, userID, roleID)
}

// GetByUserID 根据用户ID获取角色关系
func (l *UserRoleRelLogic) GetByUserID(ctx context.Context, userID int64) ([]*systype.UserRoleRelDataResp, error) {
	rels, err := l.dao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	resp := make([]*systype.UserRoleRelDataResp, 0, len(rels))
	for _, rel := range rels {
		resp = append(resp, &systype.UserRoleRelDataResp{
			ID:       rel.ID,
			TenantID: rel.TenantID,
			UserID:   rel.UserID,
			RoleID:   rel.RoleID,
		})
	}
	return resp, nil
}

// GetByRoleID 根据角色ID获取用户关系
func (l *UserRoleRelLogic) GetByRoleID(ctx context.Context, roleID int64) ([]*systype.UserRoleRelDataResp, error) {
	rels, err := l.dao.GetByRoleID(ctx, roleID)
	if err != nil {
		return nil, err
	}

	resp := make([]*systype.UserRoleRelDataResp, 0, len(rels))
	for _, rel := range rels {
		resp = append(resp, &systype.UserRoleRelDataResp{
			ID:       rel.ID,
			TenantID: rel.TenantID,
			UserID:   rel.UserID,
			RoleID:   rel.RoleID,
		})
	}
	return resp, nil
}

// DeleteByUserID 删除用户的所有角色关系
func (l *UserRoleRelLogic) DeleteByUserID(ctx context.Context, userID int64) error {
	return l.dao.DeleteByUserID(ctx, userID)
}

// DeleteByRoleID 删除角色的所有用户关系
func (l *UserRoleRelLogic) DeleteByRoleID(ctx context.Context, roleID int64) error {
	return l.dao.DeleteByRoleID(ctx, roleID)
}

// GetList 获取用户角色关系列表
func (l *UserRoleRelLogic) GetList(ctx context.Context, req *systype.UserRoleRelQueryReq) (*systype.UserRoleRelDataListResp, error) {
	rels, total, err := l.dao.List(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &systype.UserRoleRelDataListResp{
		Total: total,
		List:  make([]systype.UserRoleRelDataResp, 0, len(rels)),
	}

	for _, rel := range rels {
		resp.List = append(resp.List, systype.UserRoleRelDataResp{
			ID:       rel.ID,
			TenantID: rel.TenantID,
			UserID:   rel.UserID,
			RoleID:   rel.RoleID,
		})
	}

	return resp, nil
}