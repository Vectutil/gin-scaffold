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
func (l *UserRoleRelLogic) Create(ctx context.Context, req *systype.UserRoleRelCreateReq, tenantId int64) error {
	rel := &sysmodel.UserRoleRel{
		UserId: req.UserId,
		RoleId: req.RoleId,
	}
	rel.TenantId = tenantId
	return l.dao.Create(ctx, rel)
}

// Delete 删除用户角色关系
func (l *UserRoleRelLogic) Delete(ctx context.Context, userId, roleId int64) error {
	return l.dao.Delete(ctx, userId, roleId)
}

// GetByUserId 根据用户Id获取角色关系
func (l *UserRoleRelLogic) GetByUserId(ctx context.Context, userId int64) ([]*systype.UserRoleRelDataResp, error) {
	rels, err := l.dao.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	resp := make([]*systype.UserRoleRelDataResp, 0, len(rels))
	for _, rel := range rels {
		resp = append(resp, &systype.UserRoleRelDataResp{
			Id:       rel.Id,
			TenantId: rel.TenantId,
			UserId:   rel.UserId,
			RoleId:   rel.RoleId,
		})
	}
	return resp, nil
}

// GetByRoleId 根据角色Id获取用户关系
func (l *UserRoleRelLogic) GetByRoleId(ctx context.Context, roleId int64) ([]*systype.UserRoleRelDataResp, error) {
	rels, err := l.dao.GetByRoleId(ctx, roleId)
	if err != nil {
		return nil, err
	}

	resp := make([]*systype.UserRoleRelDataResp, 0, len(rels))
	for _, rel := range rels {
		resp = append(resp, &systype.UserRoleRelDataResp{
			Id:       rel.Id,
			TenantId: rel.TenantId,
			UserId:   rel.UserId,
			RoleId:   rel.RoleId,
		})
	}
	return resp, nil
}

// DeleteByUserId 删除用户的所有角色关系
func (l *UserRoleRelLogic) DeleteByUserId(ctx context.Context, userId int64) error {
	return l.dao.DeleteByUserId(ctx, userId)
}

// DeleteByRoleId 删除角色的所有用户关系
func (l *UserRoleRelLogic) DeleteByRoleId(ctx context.Context, roleId int64) error {
	return l.dao.DeleteByRoleId(ctx, roleId)
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
			Id:       rel.Id,
			TenantId: rel.TenantId,
			UserId:   rel.UserId,
			RoleId:   rel.RoleId,
		})
	}

	return resp, nil
}
