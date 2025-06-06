package system

import (
	"context"
	systype "gin-scaffold/internal/app/types/system"
	"time"

	sysmodel "gin-scaffold/internal/app/model/system"
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
	return d.db.WithContext(ctx).Model(user).Updates(user).Error
}

// Delete 删除用户
func (d *UserDao) Delete(ctx context.Context, id int64, deletedBy int64) error {
	return d.db.WithContext(ctx).Model(&sysmodel.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": time.Now(),
			"deleted_by": deletedBy,
		}).Error
}

// GetByID 根据ID获取用户
func (d *UserDao) GetByID(ctx context.Context, id int64) (*sysmodel.User, error) {
	var user sysmodel.User
	err := d.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername 根据用户名获取用户
func (d *UserDao) GetByUsername(ctx context.Context, username string) (*sysmodel.User, error) {
	var user sysmodel.User
	err := d.db.WithContext(ctx).Where("username = ? AND deleted_at IS NULL", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// List 查询用户列表
func (d *UserDao) List(ctx context.Context, req *systype.UserQueryReq) ([]*sysmodel.User, int64, error) {
	var users []*sysmodel.User
	var total int64

	query := d.db.WithContext(ctx).Model(&sysmodel.User{})

	if req.Username != "" {
		query = query.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.Email != "" {
		query = query.Where("email LIKE ?", "%"+req.Email+"%")
	}
	if req.Phone != "" {
		query = query.Where("phone LIKE ?", "%"+req.Phone+"%")
	}
	if req.DeptID > 0 {
		query = query.Where("dept_id = ?", req.DeptID)
	}
	if req.Status > 0 {
		query = query.Where("status = ?", req.Status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if req.Page > 0 && req.PageSize > 0 {
		query = query.Offset(req.GetOffset()).Limit(req.PageSize)
	}

	err = query.Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
