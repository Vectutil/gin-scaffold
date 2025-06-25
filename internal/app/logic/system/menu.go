package system

import (
	"context"
	"errors"
	sysdao "gin-scaffold/internal/app/dao/system"
	sysmodel "gin-scaffold/internal/app/model/system"
	"gorm.io/gorm"
)

// MenuLogic 菜单业务逻辑
type (
	MenuLogic struct {
		menuDao *sysdao.MenuDao
	}
	IMenuLogic interface {
		Create(ctx context.Context, req *sysmodel.Menu) error
		Update(ctx context.Context, req *sysmodel.Menu) error
		Delete(ctx context.Context, id int64) error
		GetById(ctx context.Context, id int64) (*sysmodel.Menu, error)
		GetList(ctx context.Context, req interface{}) ([]*sysmodel.Menu, int64, error)
	}
)

// NewMenuLogic 创建菜单Logic实例
func NewMenuLogic(db *gorm.DB) *MenuLogic {
	return &MenuLogic{
		menuDao: sysdao.NewMenuDao(db),
	}
}

// Create 创建菜单
func (l *MenuLogic) Create(ctx context.Context, req *sysmodel.Menu) error {
	// 检查菜单名称是否已存在
	existMenu, err := l.menuDao.GetById(ctx, req.Id)
	if err == nil && existMenu != nil {
		return errors.New("菜单已存在")
	}

	return l.menuDao.Create(ctx, req)
}

// Update 更新菜单
func (l *MenuLogic) Update(ctx context.Context, req *sysmodel.Menu) error {
	menu, err := l.menuDao.GetById(ctx, req.Id)
	if err != nil {
		return err
	}

	menu.Name = req.Name
	menu.Level = req.Level
	menu.Type = req.Type
	menu.Path = req.Path
	menu.PermissionKey = req.PermissionKey
	menu.OrderNum = req.OrderNum
	menu.Visible = req.Visible

	return l.menuDao.Update(ctx, menu)
}

// Delete 删除菜单
func (l *MenuLogic) Delete(ctx context.Context, id int64) error {
	menu, err := l.menuDao.GetById(ctx, id)
	if err != nil {
		return err
	}

	return l.menuDao.Delete(ctx, menu.Id)
}

// GetById 根据Id获取菜单
func (l *MenuLogic) GetById(ctx context.Context, id int64) (*sysmodel.Menu, error) {
	return l.menuDao.GetById(ctx, id)
}

// GetList 获取菜单列表
func (l *MenuLogic) GetList(ctx context.Context, req interface{}) ([]*sysmodel.Menu, int64, error) {
	return l.menuDao.List(ctx, req)
}

// GetMenuTree 获取菜单树形结构
func (l *MenuLogic) GetMenuTree(ctx context.Context) ([]*sysmodel.Menu, error) {
	allMenus, _, err := l.GetList(ctx, nil)
	if err != nil {
		return nil, err
	}

	menuMap := make(map[int64]*sysmodel.Menu)
	var rootMenus []*sysmodel.Menu

	// 构建菜单映射
	for _, menu := range allMenus {
		menuMap[menu.Id] = menu
		menu.Children = []*sysmodel.Menu{}
	}

	// 构建树形结构
	for _, menu := range allMenus {
		if menu.ParentId == 0 {
			rootMenus = append(rootMenus, menu)
		} else {
			if parent, ok := menuMap[menu.ParentId]; ok {
				parent.Children = append(parent.Children, menu)
			}
		}
	}

	return rootMenus, nil
}
