package system

import (
	"context"
	"gin-scaffold/internal/app/model/common"
	"gin-scaffold/internal/app/model/system"
	systype "gin-scaffold/internal/app/types/system"
	"gorm.io/gorm"
)

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao(db *gorm.DB) *RoleDao {
	return &RoleDao{db: db}
}

// Create 创建角色
func (d *RoleDao) Create(ctx context.Context, role *system.Role) error {
	return d.db.WithContext(ctx).Create(role).Error
}

// Update 更新角色
func (d *RoleDao) Update(ctx context.Context, role *system.Role) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Model(&system.Role{}).Where("id = ?", role.ID).Updates(role).Error
}

// Delete 删除角色
func (d *RoleDao) Delete(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Model(&system.Role{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("NOW()")).Error
}

// GetByID 根据ID获取角色
func (d *RoleDao) GetByID(ctx context.Context, id int64) (*system.Role, error) {
	var role system.Role
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// GetByCode 根据编码获取角色
func (d *RoleDao) GetByCode(ctx context.Context, code string) (*system.Role, error) {
	var role system.Role
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Where("code = ?", code).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// List 获取角色列表
func (d *RoleDao) List(ctx context.Context, req *systype.RoleQueryReq) ([]*system.Role, int64, error) {
	var (
		roles []*system.Role
		total int64
	)

	query := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Model(&system.Role{})

	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Code != "" {
		query = query.Where("code LIKE ?", "%"+req.Code+"%")
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if err := query.Offset(req.GetOffset()).Limit(req.PageSize).Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

// GetAll 获取所有角色
func (d *RoleDao) GetAll(ctx context.Context) ([]*system.Role, error) {
	var roles []*system.Role
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Find(&roles).Error
	return roles, err
}
