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
func (d *UserRoleRelDao) Delete(ctx context.Context, userID, roleID int64) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Delete(&system.UserRoleRel{}).Error
}

// GetByUserID 根据用户ID获取角色关系
func (d *UserRoleRelDao) GetByUserID(ctx context.Context, userID int64) ([]*system.UserRoleRel, error) {
	var rels []*system.UserRoleRel
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Where("user_id = ?", userID).Find(&rels).Error
	return rels, err
}

// GetByRoleID 根据角色ID获取用户关系
func (d *UserRoleRelDao) GetByRoleID(ctx context.Context, roleID int64) ([]*system.UserRoleRel, error) {
	var rels []*system.UserRoleRel
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Where("role_id = ?", roleID).Find(&rels).Error
	return rels, err
}

// DeleteByUserID 删除用户的所有角色关系
func (d *UserRoleRelDao) DeleteByUserID(ctx context.Context, userID int64) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Where("user_id = ?", userID).
		Delete(&system.UserRoleRel{}).Error
}

// DeleteByRoleID 删除角色的所有用户关系
func (d *UserRoleRelDao) DeleteByRoleID(ctx context.Context, roleID int64) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Where("role_id = ?", roleID).
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

	if req.UserID != 0 {
		query = query.Where("user_id = ?", req.UserID)
	}
	if req.RoleID != 0 {
		query = query.Where("role_id = ?", req.RoleID)
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