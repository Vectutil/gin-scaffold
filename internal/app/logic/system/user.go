package system

import (
	"context"
	"errors"
	sysdao "gin-scaffold/internal/app/dao/system"
	sysmodel "gin-scaffold/internal/app/model/system"
	"gin-scaffold/internal/app/types/common"
	systype "gin-scaffold/internal/app/types/system"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

// userLogic 用户业务逻辑
type (
	userLogic struct {
		userDao *sysdao.UserDao
	}
	IUserLogic interface {
		Create(ctx context.Context, req *systype.UserCreateReq, operatorID int64) error
		Update(ctx context.Context, req *systype.UserUpdateReq, operatorID int64) error
		Delete(ctx context.Context, id int64, operatorID int64) error
		GetByID(ctx context.Context, id int64) (*systype.UserDataResp, error)
		UpdateLoginInfo(ctx context.Context, id int64, ip string) error
		GetList(ctx context.Context, req *systype.UserQueryReq) (*systype.UserDataListResp, error)
	}
)

// NewUserLogic 创建用户Logic实例
func NewUserLogic(db *gorm.DB) *userLogic {
	return &userLogic{userDao: sysdao.NewUserDao(db)}
}

// Create 创建用户
func (l *userLogic) Create(ctx context.Context, req *systype.UserCreateReq, operatorID int64) error {
	// 检查用户名是否已存在
	existUser, err := l.userDao.GetByUsername(ctx, req.Username)
	if err == nil && existUser != nil {
		return errors.New("用户名已存在")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &sysmodel.User{
		Username:  req.Username,
		Password:  string(hashedPassword),
		FullName:  req.FullName,
		Email:     req.Email,
		Phone:     req.Phone,
		DeptID:    req.DeptID,
		Status:    req.Status,
		Remark:    req.Remark,
		CreatedBy: operatorID,
		UpdatedBy: operatorID,
	}

	return l.userDao.Create(ctx, user)
}

// Update 更新用户
func (l *userLogic) Update(ctx context.Context, req *systype.UserUpdateReq, operatorID int64) error {
	user, err := l.userDao.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}

	user.FullName = req.FullName
	user.Email = req.Email
	user.Phone = req.Phone
	user.DeptID = req.DeptID
	user.Status = req.Status
	user.Remark = req.Remark
	user.UpdatedBy = operatorID

	return l.userDao.Update(ctx, user)
}

// UpdateLoginInfo 更新登录信息
func (l *userLogic) UpdateLoginInfo(ctx context.Context, id int64, ip string) error {
	user, err := l.userDao.GetByID(ctx, id)
	if err != nil {
		return err
	}

	user.LoginCount++
	user.LastLoginAt = time.Now().Unix()
	user.LastLoginIP = ip

	return l.userDao.Update(ctx, user)
}

// Delete 删除用户
func (l *userLogic) Delete(ctx context.Context, id int64, operatorID int64) error {
	user, err := l.userDao.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return l.userDao.Delete(ctx, user.ID, operatorID)
}

// GetByID 根据ID获取用户
func (l *userLogic) GetByID(ctx context.Context, id int64) (*systype.UserDataResp, error) {
	user, err := l.userDao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &systype.UserDataResp{
		ID:          user.ID,
		Username:    user.Username,
		FullName:    user.FullName,
		Avatar:      user.Avatar,
		Email:       user.Email,
		Phone:       user.Phone,
		DeptID:      user.DeptID,
		Status:      user.Status,
		LoginCount:  user.LoginCount,
		LastLoginAt: user.LastLoginAt,
		LastLoginIP: user.LastLoginIP,
		TenantID:    user.TenantID,
		OrgID:       user.OrgID,
		Remark:      user.Remark,
		CreatedAt:   user.CreatedAt,
		CreatedBy:   user.CreatedBy,
		UpdatedAt:   user.UpdatedAt,
		UpdatedBy:   user.UpdatedBy,
	}, nil
}

// GetList 查询用户列表
func (l *userLogic) GetList(ctx context.Context, req *systype.UserQueryReq) (*systype.UserDataListResp, error) {
	users, total, err := l.userDao.List(ctx, req)
	if err != nil {
		return nil, err
	}

	list := make([]*systype.UserDataResp, 0, len(users))
	for _, user := range users {
		list = append(list, &systype.UserDataResp{
			ID:          user.ID,
			Username:    user.Username,
			FullName:    user.FullName,
			Avatar:      user.Avatar,
			Email:       user.Email,
			Phone:       user.Phone,
			DeptID:      user.DeptID,
			Status:      user.Status,
			LoginCount:  user.LoginCount,
			LastLoginAt: user.LastLoginAt,
			LastLoginIP: user.LastLoginIP,
			TenantID:    user.TenantID,
			OrgID:       user.OrgID,
			Remark:      user.Remark,
			CreatedAt:   user.CreatedAt,
			CreatedBy:   user.CreatedBy,
			UpdatedAt:   user.UpdatedAt,
			UpdatedBy:   user.UpdatedBy,
		})
	}
	res := &systype.UserDataListResp{
		List: list,
		ListResp: common.ListResp{
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}
	res.TotalPage = res.GetTotalPage()
	return res, nil
}
