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
		Name:     req.DeptName,
		ParentID: *req.ParentID,
		Status:   req.Status,
	}

	return l.dao.Create(ctx, dept)
}

// Update 更新部门
func (l *DepartmentLogic) Update(ctx context.Context, req *systype.DepartmentUpdateReq, operatorID int64) error {
	// 检查部门是否存在
	dept, err := l.dao.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}

	dept.Name = req.DeptName
	dept.ParentID = *req.ParentID
	dept.Status = req.Status

	return l.dao.Update(ctx, dept)
}

// Delete 删除部门
func (l *DepartmentLogic) Delete(ctx context.Context, id int64, operatorID int64) error {
	// 检查是否存在子部门
	count, err := l.dao.CountByParentID(ctx, id)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("存在子部门，无法删除")
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
		DeptName:  dept.Name,
		TenantID:  dept.TenantID,
		ParentID:  &dept.ParentID,
		Status:    dept.Status,
		CreatedAt: dept.CreatedAt,
		CreatedBy: dept.CreatedBy,
		UpdatedAt: dept.UpdatedAt,
		UpdatedBy: dept.UpdatedBy,
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
			DeptName:  dept.Name,
			TenantID:  dept.TenantID,
			ParentID:  &dept.ParentID,
			Status:    dept.Status,
			CreatedAt: dept.CreatedAt,
			CreatedBy: dept.CreatedBy,
			UpdatedAt: dept.UpdatedAt,
			UpdatedBy: dept.UpdatedBy,
		})
	}

	return resp, nil
}

// GetTree 获取部门树
func (l *DepartmentLogic) GetTree(ctx context.Context, tenantID int64) ([]systype.DepartmentTreeResp, error) {
	trees, err := l.dao.GetTree(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	resp := make([]systype.DepartmentTreeResp, 0, len(trees))
	for _, tree := range trees {
		resp = append(resp, systype.DepartmentTreeResp{
			ID:       tree.ID,
			DeptName: tree.Name,
			TenantID: tree.TenantID,
			ParentID: tree.ParentID,
			Status:   tree.Status,
			Children: convertTreeToResp(tree.Children),
		})
	}

	return resp, nil
}

func convertTreeToResp(trees []sysmodel.DepartmentTree) []systype.DepartmentTreeResp {
	resp := make([]systype.DepartmentTreeResp, 0, len(trees))
	for _, tree := range trees {
		resp = append(resp, systype.DepartmentTreeResp{
			ID:       tree.ID,
			DeptName: tree.Name,
			TenantID: tree.TenantID,
			ParentID: tree.ParentID,
			Status:   tree.Status,
			Children: convertTreeToResp(tree.Children),
		})
	}
	return resp
}
