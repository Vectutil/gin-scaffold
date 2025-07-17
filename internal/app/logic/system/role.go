package system

import (
	"context"
	"errors"
	sysdao "gin-scaffold/internal/app/dao/system"
	sysmodel "gin-scaffold/internal/app/model/system"
	systype "gin-scaffold/internal/app/types/system"
	"gorm.io/gorm"
)

// RoleLogic 角色逻辑
type RoleLogic struct {
	dao *sysdao.RoleDao
}

// NewRoleLogic 创建角色逻辑实例
func NewRoleLogic(db *gorm.DB) *RoleLogic {
	return &RoleLogic{dao: sysdao.NewRoleDao(db)}
}

// Create 创建角色
func (l *RoleLogic) Create(ctx context.Context, req *systype.RoleCreateReq) error {
	// 检查角色编码是否已存在
	_, err := l.dao.GetByCode(ctx, req.Code)
	if err == nil {
		return errors.New("角色编码已存在")
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}

	role := &sysmodel.Role{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
		DataScope:   req.DataScope,
		Status:      req.Status,
	}

	return l.dao.Create(ctx, role)
}

// Update 更新角色
func (l *RoleLogic) Update(ctx context.Context, req *systype.RoleUpdateReq) error {
	// 检查角色是否存在
	role, err := l.dao.GetById(ctx, req.Id)
	if err != nil {
		return err
	}

	// 检查角色编码是否已被其他角色使用
	if role.Code != req.Code {
		existingRole, err := l.dao.GetByCode(ctx, req.Code)
		if err == nil && existingRole.Id != req.Id {
			return errors.New("角色编码已被其他角色使用")
		}
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
	}

	role.Name = req.Name
	role.Code = req.Code
	role.Description = req.Description
	role.DataScope = req.DataScope
	role.Status = req.Status

	return l.dao.Update(ctx, role)
}

// Delete 删除角色
func (l *RoleLogic) Delete(ctx context.Context, id int64) error {
	return l.dao.Delete(ctx, id)
}

// GetById 根据Id获取角色
func (l *RoleLogic) GetById(ctx context.Context, id int64) (*systype.RoleDataResp, error) {
	role, err := l.dao.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &systype.RoleDataResp{
		Id:          role.Id,
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
		DataScope:   role.DataScope,
		Status:      role.Status,
		CreatedAt:   role.CreatedAt,
		CreatedBy:   role.CreatedBy,
		UpdatedAt:   role.UpdatedAt,
		UpdatedBy:   role.UpdatedBy,
	}, nil
}

// GetList 获取角色列表
func (l *RoleLogic) GetList(ctx context.Context, req *systype.RoleQueryReq) (*systype.RoleDataListResp, error) {
	roles, total, err := l.dao.List(ctx, req)
	if err != nil {
		return nil, err
	}

	// 构建响应
	resp := &systype.RoleDataListResp{
		Total: total,
		List:  make([]systype.RoleDataResp, 0, len(roles)),
	}

	for _, role := range roles {
		resp.List = append(resp.List, systype.RoleDataResp{
			Id:          role.Id,
			Name:        role.Name,
			Code:        role.Code,
			Description: role.Description,
			DataScope:   role.DataScope,
			Status:      role.Status,
			CreatedAt:   role.CreatedAt,
			CreatedBy:   role.CreatedBy,
			UpdatedAt:   role.UpdatedAt,
			UpdatedBy:   role.UpdatedBy,
		})
	}

	return resp, nil
}
