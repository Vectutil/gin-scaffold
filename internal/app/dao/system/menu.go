package system

import (
	"context"
	"gin-scaffold/internal/app/model/common"
	"gin-scaffold/internal/app/model/system"
	"gorm.io/gorm"
)

type MenuDao struct {
	db *gorm.DB
}

func NewMenuDao(db *gorm.DB) *MenuDao {
	return &MenuDao{db: db}
}

// Create 创建菜单
func (d *MenuDao) Create(ctx context.Context, menu *system.Menu) error {
	return d.db.WithContext(ctx).Create(menu).Error
}

// Update 更新菜单
func (d *MenuDao) Update(ctx context.Context, menu *system.Menu) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx), common.NotDeletedScope()).
		Model(&system.Menu{}).Where("id = ?", menu.ID).Updates(menu).Error
}

// Delete 删除菜单
func (d *MenuDao) Delete(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx), common.NotDeletedScope()).
		Model(&system.Menu{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("NOW()")).Error
}

// GetByID 根据ID获取菜单
func (d *MenuDao) GetByID(ctx context.Context, id int64) (*system.Menu, error) {
	var menu system.Menu
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx), common.NotDeletedScope()).
		First(&menu, id).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

// List 获取菜单列表
func (d *MenuDao) List(ctx context.Context, req interface{}) ([]*system.Menu, int64, error) {
	var (
		menus []*system.Menu
		total int64
	)

	query := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx), common.NotDeletedScope()).
		Model(&system.Menu{})

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询数据
	if err := query.Find(&menus).Error; err != nil {
		return nil, 0, err
	}

	return menus, total, nil
}