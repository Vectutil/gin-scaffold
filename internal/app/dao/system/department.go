package system

import (
	"context"
	"gin-scaffold/internal/app/model/common"
	"gin-scaffold/internal/app/model/system"
	systype "gin-scaffold/internal/app/types/system"
	"gorm.io/gorm"
)

type DepartmentDao struct {
	db *gorm.DB
}

func NewDepartmentDao(db *gorm.DB) *DepartmentDao {
	return &DepartmentDao{db: db}
}

// Create 创建部门
func (d *DepartmentDao) Create(ctx context.Context, dept *system.Department) error {
	return d.db.WithContext(ctx).Create(dept).Error
}

// Update 更新部门
func (d *DepartmentDao) Update(ctx context.Context, dept *system.Department) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx), common.NotDeletedScope()).
		Model(&system.Department{}).Where("id = ?", dept.ID).Updates(dept).Error
}

// Delete 删除部门
func (d *DepartmentDao) Delete(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx), common.NotDeletedScope()).
		Model(&system.Department{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("NOW()")).Error
}

// GetByID 根据ID获取部门
func (d *DepartmentDao) GetByID(ctx context.Context, id int64) (*system.Department, error) {
	var dept system.Department
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx), common.NotDeletedScope()).
		First(&dept, id).Error
	if err != nil {
		return nil, err
	}
	return &dept, nil
}

// CountByParentID 统计子部门数量
func (d *DepartmentDao) CountByParentID(ctx context.Context, parentID int64) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx), common.NotDeletedScope()).
		Model(&system.Department{}).Where("parent_id = ?", parentID).Count(&count).Error
	return count, err
}

// List 获取部门列表
func (d *DepartmentDao) List(ctx context.Context, req *systype.DepartmentQueryReq) ([]*system.Department, int64, error) {
	var (
		depts []*system.Department
		total int64
	)

	query := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx), common.NotDeletedScope()).
		Model(&system.Department{})

	if req.DeptName != "" {
		query = query.Where("dept_name LIKE ?", "%"+req.DeptName+"%")
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if err := query.Offset(req.GetOffset()).Limit(req.PageSize).Find(&depts).Error; err != nil {
		return nil, 0, err
	}

	return depts, total, nil
}

// GetAll 获取所有部门
func (d *DepartmentDao) GetAll(ctx context.Context) ([]*system.Department, error) {
	var depts []*system.Department
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx), common.NotDeletedScope()).
		Find(&depts).Error
	return depts, err
}

// GetChildren 获取子部门
func (d *DepartmentDao) GetChildren(ctx context.Context, parentID int64) ([]*system.Department, error) {
	var depts []*system.Department
	err := d.db.WithContext(ctx).Where("parent_id = ?", parentID).Find(&depts).Error
	if err != nil {
		return nil, err
	}
	return depts, nil
}

// GetTree 获取部门树
func (d *DepartmentDao) GetTree(ctx context.Context, tenantID int64) ([]system.DepartmentTree, error) {
	var depts []*system.Department
	err := d.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Find(&depts).Error
	if err != nil {
		return nil, err
	}

	// 构建部门树
	deptMap := make(map[int64]*system.DepartmentTree)
	var roots []system.DepartmentTree

	// 先将所有部门转换为树节点
	for _, dept := range depts {
		tree := system.DepartmentTree{
			ID:       dept.ID,
			DeptName: dept.DeptName,
			TenantID: dept.TenantID,
			ParentID: dept.ParentID,
			Status:   dept.Status,
		}
		deptMap[dept.ID] = &tree
	}

	// 构建树结构
	for _, dept := range depts {
		if dept.ParentID == nil {
			// 根节点
			roots = append(roots, *deptMap[dept.ID])
		} else {
			// 子节点
			parent := deptMap[*dept.ParentID]
			parent.Children = append(parent.Children, *deptMap[dept.ID])
		}
	}

	return roots, nil
}
