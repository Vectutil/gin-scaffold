package system

import (
	"context"
	"gin-scaffold/internal/app/model/common"
	"gin-scaffold/internal/app/model/system"
	systype "gin-scaffold/internal/app/types/system"
	"gorm.io/gorm"
)

type UserRoleRelDao struct {
	db *gorm.DB
}

func NewUserRoleRelDao(db *gorm.DB) *UserRoleRelDao {
	return &UserRoleRelDao{db: db}
}

// Create 创建用户角色关系
func (d *UserRoleRelDao) Create(ctx context.Context, rel *system.UserRoleRel) error {
	return d.db.WithContext(ctx).Create(rel).Error
}

// CreateList 批量创建用户角色关系
func (d *UserRoleRelDao) CreateList(ctx context.Context, urList []system.UserRoleRel) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).Create(&urList).Error
}

// Delete 删除用户角色关系
func (d *UserRoleRelDao) Delete(ctx context.Context, userId, roleId int64) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Where("user_id = ? AND role_id = ?", userId, roleId).
		Delete(&system.UserRoleRel{}).Error
}

// GetByUserId 根据用户Id获取角色关系
func (d *UserRoleRelDao) GetByUserId(ctx context.Context, userId int64) ([]*system.UserRoleRel, error) {
	var rels []*system.UserRoleRel
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Where("user_id = ?", userId).Find(&rels).Error
	return rels, err
}

// GetByRoleId 根据角色Id获取用户关系
func (d *UserRoleRelDao) GetByRoleId(ctx context.Context, roleId int64) ([]*system.UserRoleRel, error) {
	var rels []*system.UserRoleRel
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Where("role_id = ?", roleId).Find(&rels).Error
	return rels, err
}

// DeleteByUserId 删除用户的所有角色关系
func (d *UserRoleRelDao) DeleteByUserId(ctx context.Context, userId int64) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Where("user_id = ?", userId).
		Delete(&system.UserRoleRel{}).Error
}

// DeleteByRoleId 删除角色的所有用户关系
func (d *UserRoleRelDao) DeleteByRoleId(ctx context.Context, roleId int64) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Where("role_id = ?", roleId).
		Delete(&system.UserRoleRel{}).Error
}

// List 获取用户角色关系列表
func (d *UserRoleRelDao) List(ctx context.Context, req *systype.UserRoleRelQueryReq) ([]*system.UserRoleRel, int64, error) {
	var (
		rels  []*system.UserRoleRel
		total int64
	)

	query := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Model(&system.UserRoleRel{})

	if req.UserId != 0 {
		query = query.Where("user_id = ?", req.UserId)
	}
	if req.RoleId != 0 {
		query = query.Where("role_id = ?", req.RoleId)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询数据
	if err := query.Find(&rels).Error; err != nil {
		return nil, 0, err
	}

	return rels, total, nil
}
