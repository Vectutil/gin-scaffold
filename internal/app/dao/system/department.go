package system

import (
	"context"
	"gin-scaffold/internal/app/model/common"
	sysmodel "gin-scaffold/internal/app/model/system"
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
func (d *DepartmentDao) Create(ctx context.Context, dept *sysmodel.Department) error {
	return d.db.WithContext(ctx).Create(dept).Error
}

// Update 更新部门
func (d *DepartmentDao) Update(ctx context.Context, dept *sysmodel.Department) error {
	return d.db.WithContext(ctx).
		Model(&sysmodel.Department{}).Where("id = ?", dept.Id).Save(dept).Error
}

// Delete 删除部门
func (d *DepartmentDao) Delete(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Delete(&sysmodel.Department{
		BaseModel: common.BaseModel{
			Id: id,
		},
	}).Error
}

// GetById 根据Id获取部门
func (d *DepartmentDao) GetById(ctx context.Context, id int64) (*sysmodel.Department, error) {
	var dept sysmodel.Department
	err := d.db.WithContext(ctx).
		First(&dept, id).Error
	if err != nil {
		return nil, err
	}
	return &dept, nil
}

// CountByParentId 统计子部门数量
func (d *DepartmentDao) CountByParentId(ctx context.Context, parentId int64) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).
		Model(&sysmodel.Department{}).Where("parent_id = ?", parentId).Count(&count).Error
	return count, err
}

// List 获取部门列表
func (d *DepartmentDao) List(ctx context.Context, req *systype.DepartmentQueryReq) ([]*sysmodel.Department, int64, error) {
	var (
		depts []*sysmodel.Department
		total int64
	)

	query := d.db.WithContext(ctx).
		Model(&sysmodel.Department{})

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
func (d *DepartmentDao) GetAll(ctx context.Context) ([]*sysmodel.Department, error) {
	var depts []*sysmodel.Department
	err := d.db.WithContext(ctx).
		Find(&depts).Error
	return depts, err
}

// GetChildren 获取子部门
func (d *DepartmentDao) GetChildren(ctx context.Context, parentId int64) ([]*sysmodel.Department, error) {
	var depts []*sysmodel.Department
	err := d.db.WithContext(ctx).Where("parent_id = ?", parentId).Find(&depts).Error
	if err != nil {
		return nil, err
	}
	return depts, nil
}

// GetTree 获取部门树
func (d *DepartmentDao) GetTree(ctx context.Context) ([]*sysmodel.DepartmentTree, error) {
	var depts []*sysmodel.Department
	err := d.db.WithContext(ctx).Find(&depts).Error
	if err != nil {
		return nil, err
	}

	deptMap := make(map[int64]*sysmodel.DepartmentTree)
	var roots []*sysmodel.DepartmentTree

	// 先构造所有节点
	for _, dept := range depts {
		node := &sysmodel.DepartmentTree{
			Name:     dept.Name,
			ParentId: dept.ParentId,
			Status:   dept.Status,
			Children: []*sysmodel.DepartmentTree{},
		}
		node.Id = dept.Id
		deptMap[dept.Id] = node
	}

	// 构建树
	for _, node := range deptMap {
		if node.ParentId == 0 {
			roots = append(roots, node)
		} else if parent, ok := deptMap[node.ParentId]; ok {
			parent.Children = append(parent.Children, node) // 用指针
		}
	}

	return roots, nil
}
