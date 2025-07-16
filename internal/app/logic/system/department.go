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
func (l *DepartmentLogic) Create(ctx context.Context, req *systype.DepartmentCreateReq, operatorId int64) error {
	dept := &sysmodel.Department{
		Name:     req.DeptName,
		ParentId: *req.ParentId,
		Status:   req.Status,
	}

	return l.dao.Create(ctx, dept)
}

// Update 更新部门
func (l *DepartmentLogic) Update(ctx context.Context, req *systype.DepartmentUpdateReq, operatorId int64) error {
	// 检查部门是否存在
	dept, err := l.dao.GetById(ctx, req.Id)
	if err != nil {
		return err
	}

	dept.Name = req.DeptName
	dept.ParentId = *req.ParentId
	dept.Status = req.Status

	return l.dao.Update(ctx, dept)
}

// Delete 删除部门
func (l *DepartmentLogic) Delete(ctx context.Context, id int64, operatorId int64) error {
	// 检查是否存在子部门
	count, err := l.dao.CountByParentId(ctx, id)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("存在子部门，无法删除")
	}

	return l.dao.Delete(ctx, id)
}

// GetById 根据Id获取部门
func (l *DepartmentLogic) GetById(ctx context.Context, id int64) (*systype.DepartmentDataResp, error) {
	dept, err := l.dao.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &systype.DepartmentDataResp{
		Id:       dept.Id,
		DeptName: dept.Name,
		//TenantId:  dept.TenantId,
		ParentId:  &dept.ParentId,
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
			Id:       dept.Id,
			DeptName: dept.Name,
			//TenantId:  dept.TenantId,
			ParentId:  &dept.ParentId,
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
func (l *DepartmentLogic) GetTree(ctx context.Context) ([]systype.DepartmentTreeResp, error) {
	trees, err := l.dao.GetTree(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]systype.DepartmentTreeResp, 0, len(trees))
	for _, tree := range trees {
		resp = append(resp, systype.DepartmentTreeResp{
			Id:       tree.Id,
			DeptName: tree.Name,
			//TenantId: tree.TenantId,
			ParentId: tree.ParentId,
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
			Id:       tree.Id,
			DeptName: tree.Name,
			//TenantId: tree.TenantId,
			ParentId: tree.ParentId,
			Status:   tree.Status,
			Children: convertTreeToResp(tree.Children),
		})
	}
	return resp
}
