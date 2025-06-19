package system

import (
	"context"
	"gin-scaffold/internal/app/model/common"
	sysmodel "gin-scaffold/internal/app/model/system"
	systype "gin-scaffold/internal/app/types/system"

	"gorm.io/gorm"
)

// UserDao 用户数据访问对象
type (
	UserDao struct {
		db *gorm.DB
	}
)

// NewUserDao 创建用户DAO实例
func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// Create 创建用户
func (d *UserDao) Create(ctx context.Context, user *sysmodel.User) error {
	return d.db.WithContext(ctx).Create(user).Error
}

// Update 更新用户
func (d *UserDao) Update(ctx context.Context, user *sysmodel.User) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Model(&sysmodel.User{}).Where("id = ?", user.ID).Updates(user).Error
}

// Delete 删除用户
func (d *UserDao) Delete(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Model(&sysmodel.User{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("NOW()")).Error
}

// GetByID 根据ID获取用户
func (d *UserDao) GetByID(ctx context.Context, id int64) (*sysmodel.User, error) {
	var user sysmodel.User
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername 根据用户名获取用户
func (d *UserDao) GetByUsername(ctx context.Context, username string) (*sysmodel.User, error) {
	var user sysmodel.User
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByPhone 根据手机号获取用户
func (d *UserDao) GetByPhone(ctx context.Context, phone string) (*sysmodel.User, error) {
	var user sysmodel.User
	err := d.db.WithContext(ctx).Where("phone = ? AND deleted_at IS NULL", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// List 查询用户列表
func (d *UserDao) List(ctx context.Context, req *systype.UserQueryReq) ([]*sysmodel.User, int64, error) {
	var (
		users []*sysmodel.User
		total int64
	)

	query := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Model(&sysmodel.User{})

	// 添加查询条件
	if req.Username != "" {
		query = query.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.Status != 0 {
		query = query.Where("status = ?", req.Status)
	}
	if req.DeptID != 0 {
		query = query.Where("dept_id = ?", req.DeptID)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if err := query.Offset(req.GetOffset()).Limit(req.PageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// UpdatePassword 更新密码
func (d *UserDao) UpdatePassword(ctx context.Context, id int64, password string) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Model(&sysmodel.User{}).Where("id = ?", id).Update("password", password).Error
}

// UpdateStatus 更新状态
func (d *UserDao) UpdateStatus(ctx context.Context, id int64, status int) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Model(&sysmodel.User{}).Where("id = ?", id).Update("status", status).Error
}

// UpdateDept 更新部门
func (d *UserDao) UpdateDept(ctx context.Context, id int64, deptID int64) error {
	return d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Model(&sysmodel.User{}).Where("id = ?", id).Update("dept_id", deptID).Error
}

// CountByDeptID 统计部门下的用户数量
func (d *UserDao) CountByDeptID(ctx context.Context, deptID int64) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).Scopes(common.TenantScope(ctx)).
		Model(&sysmodel.User{}).Where("dept_id = ?", deptID).Count(&count).Error
	return count, err
}
