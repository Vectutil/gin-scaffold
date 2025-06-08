package system

import (
	"context"
	"errors"
	sysdao "gin-scaffold/internal/app/dao/system"
	sysmodel "gin-scaffold/internal/app/model/system"
	systype "gin-scaffold/internal/app/types/system"
	"gorm.io/gorm"
)

// DepartmentLogic 部门逻辑
type DepartmentLogic struct {
	dao *sysdao.DepartmentDao
}

// NewDepartmentLogic 创建部门逻辑实例
func NewDepartmentLogic(db *gorm.DB) *DepartmentLogic {
	return &DepartmentLogic{dao: sysdao.NewDepartmentDao(db)}
}

// Create 创建部门
func (l *DepartmentLogic) Create(ctx context.Context, req *systype.DepartmentCreateReq, operatorID int64) error {
	dept := &sysmodel.Department{
		DeptName: req.DeptName,
		ParentID: req.ParentID,
		Status:   req.Status,
	}

	return l.dao.Create(ctx, dept)
}

// Update 更新部门
func (l *DepartmentLogic) Update(ctx context.Context, req *systype.DepartmentUpdateReq, operatorID int64) error {
	dept := &sysmodel.Department{
		DeptName: req.DeptName,
		ParentID: req.ParentID,
		Status:   req.Status,
	}
	dept.ID = req.ID

	return l.dao.Update(ctx, dept)
}

// Delete 删除部门
func (l *DepartmentLogic) Delete(ctx context.Context, id int64, operatorID int64) error {
	// 检查是否有子部门
	count, err := l.dao.CountByParentID(ctx, id)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("部门下存在子部门，无法删除")
	}

	return l.dao.Delete(ctx, id)
}

// GetByID 根据ID获取部门
func (l *DepartmentLogic) GetByID(ctx context.Context, id int64) (*systype.DepartmentDataResp, error) {
	dept, err := l.dao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &systype.DepartmentDataResp{
		ID:        dept.ID,
		DeptName:  dept.DeptName,
		TenantID:  dept.TenantID,
		ParentID:  dept.ParentID,
		Status:    dept.Status,
		CreatedAt: dept.CreatedAt,
		CreatedBy: dept.CreatedBy,
		UpdatedAt: dept.UpdatedAt,
		UpdatedBy: dept.UpdatedBy,
		DeletedAt: dept.DeletedAt,
		DeletedBy: dept.DeletedBy,
	}, nil
}

// GetList 获取部门列表
func (l *DepartmentLogic) GetList(ctx context.Context, req *systype.DepartmentQueryReq) (*systype.DepartmentDataListResp, error) {
	depts, total, err := l.dao.List(ctx, req)
	if err != nil {
		return nil, err
	}

	// 构建响应
	resp := &systype.DepartmentDataListResp{
		Total: total,
		List:  make([]systype.DepartmentDataResp, 0, len(depts)),
	}

	for _, dept := range depts {
		resp.List = append(resp.List, systype.DepartmentDataResp{
			ID:        dept.ID,
			DeptName:  dept.DeptName,
			TenantID:  dept.TenantID,
			ParentID:  dept.ParentID,
			Status:    dept.Status,
			CreatedAt: dept.CreatedAt,
			CreatedBy: dept.CreatedBy,
			UpdatedAt: dept.UpdatedAt,
			UpdatedBy: dept.UpdatedBy,
			DeletedAt: dept.DeletedAt,
			DeletedBy: dept.DeletedBy,
		})
	}

	return resp, nil
}

// GetTree 获取部门树
func (l *DepartmentLogic) GetTree(ctx context.Context, tenantID int64) ([]systype.DepartmentTreeResp, error) {
	depts, err := l.dao.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// 构建部门树
	deptMap := make(map[int64]*systype.DepartmentTreeResp)
	var roots []systype.DepartmentTreeResp

	// 先将所有部门转换为树节点
	for _, dept := range depts {
		tree := systype.DepartmentTreeResp{
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
