package system

import (
	"context"
	"errors"
	sysdao "gin-scaffold/internal/app/dao/system"
	sysmodel "gin-scaffold/internal/app/model/system"
	"gin-scaffold/internal/app/response"
	"gin-scaffold/internal/app/types/common"
	systype "gin-scaffold/internal/app/types/system"
	"gin-scaffold/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// userLogic 用户业务逻辑
type (
	userLogic struct {
		userDao     *sysdao.UserDao
		roleDao     *sysdao.RoleDao
		roleUserDao *sysdao.UserRoleRelDao
	}
	IUserLogic interface {
		Create(ctx context.Context, req *systype.UserCreateReq) error
		CreateForRegister(ctx context.Context, req *systype.UserCreateReq) error
		Update(ctx context.Context, req *systype.UserUpdateReq) error
		Delete(ctx context.Context, id int64) error
		GetById(ctx context.Context, id int64) (*systype.UserDataResp, error)
		UpdateLoginInfo(ctx context.Context, id int64, ip string) error
		GetList(ctx context.Context, req *systype.UserQueryReq) (*systype.UserDataListResp, error)
		GetByPhone(ctx context.Context, phone string) (*systype.UserDataResp, error)
		CheckForLogin(ctx context.Context, phone string, password string) (*systype.UserDataResp, error)
	}
)

// NewUserLogic 创建用户Logic实例
func NewUserLogic(db *gorm.DB) *userLogic {
	return &userLogic{
		userDao:     sysdao.NewUserDao(db),
		roleDao:     sysdao.NewRoleDao(db),
		roleUserDao: sysdao.NewUserRoleRelDao(db),
	}
}

// Create 创建用户
func (l *userLogic) Create(ctx context.Context, req *systype.UserCreateReq) error {
	// 检查用户名是否已存在
	existUser, err := l.userDao.GetByUsername(ctx, req.Username)
	if err == nil && existUser != nil {
		return errors.New("用户名已存在")
	}

	// 密码加密
	hashedPassword, err := utils.GetHashStr(req.Password)
	if err != nil {
		return err
	}

	user := &sysmodel.User{
		Username: req.Username,
		Password: string(hashedPassword),
		FullName: req.FullName,
		Email:    req.Email,
		Phone:    req.Phone,
		DeptId:   req.DeptId,
		Status:   req.Status,
		Remark:   req.Remark,
	}

	//tenantId, err := utils.GetTenantIdFromContext(ctx)
	//if err != nil {
	//	return err
	//}

	err = l.userDao.Create(ctx, user)

	urList := make([]sysmodel.UserRoleRel, 0)

	for _, id := range req.RoleIds {
		rel := sysmodel.UserRoleRel{
			UserId: user.Id,
			RoleId: id,
		}
		//rel.TenantId = tenantId
		urList = append(urList, rel)
	}

	return l.roleUserDao.CreateList(ctx, urList)
}
func (l *userLogic) CreateForRegister(ctx context.Context, req *systype.UserCreateReq) error {
	// 检查用户名是否已存在
	existUser, err := l.userDao.GetByUsername(ctx, req.Username)
	if err == nil && existUser != nil {
		return errors.New("用户名已存在")
	}

	// 密码加密
	hashedPassword, err := utils.GetHashStr(req.Password)
	if err != nil {
		return err
	}

	user := &sysmodel.User{
		Username: req.Username,
		Password: string(hashedPassword),
		FullName: req.FullName,
		Email:    req.Email,
		Phone:    req.Phone,
		DeptId:   req.DeptId,
		Status:   req.Status,
		Remark:   req.Remark,
	}

	//tenantId, err := utils.GetTenantIdFromContext(ctx)
	//if err != nil {
	//	return err
	//}

	err = l.userDao.Create(ctx, user)
	if err != nil {
		return err
	}

	urList := make([]sysmodel.UserRoleRel, 0)

	for _, id := range req.RoleIds {
		rel := sysmodel.UserRoleRel{
			UserId: user.Id,
			RoleId: id,
		}
		//rel.TenantId = tenantId
		urList = append(urList, rel)
	}

	return l.roleUserDao.CreateList(ctx, urList)
}

// Update 更新用户
func (l *userLogic) Update(ctx context.Context, req *systype.UserUpdateReq) error {
	user, err := l.userDao.GetById(ctx, req.Id)
	if err != nil {
		return err
	}
	if user.Phone != req.Phone {
		existUser, _err := l.userDao.GetByPhone(ctx, req.Phone)
		if _err == nil && existUser != nil {
			return errors.New("用户名已存在")
		}
	}

	user.FullName = req.FullName
	user.Email = req.Email
	user.Phone = req.Phone
	user.DeptId = req.DeptId
	user.Status = req.Status
	user.Remark = req.Remark

	err = l.userDao.Update(ctx, user)
	if err != nil {
		return err
	}

	// 删除用户与角色的旧关系
	//tenantId, err := utils.GetTenantIdFromContext(ctx)
	//if err != nil {
	//	return err
	//}
	err = l.roleUserDao.DeleteByUserId(ctx, req.Id)
	if err != nil {
		return err
	}

	// 插入用户与角色的新关系
	urList := make([]sysmodel.UserRoleRel, 0)
	for _, id := range req.RoleIds {
		rel := sysmodel.UserRoleRel{
			UserId: user.Id,
			RoleId: id,
		}
		//rel.TenantId = tenantId
		urList = append(urList, rel)
	}
	return l.roleUserDao.CreateList(ctx, urList)

}

// UpdateLoginInfo 更新登录信息
func (l *userLogic) UpdateLoginInfo(ctx context.Context, id int64, ip string) error {
	user, err := l.userDao.GetById(ctx, id)
	if err != nil {
		return err
	}

	user.LoginCount++
	user.LastLoginAt = time.Now().Unix()
	user.LastLoginIP = ip

	return l.userDao.Update(ctx, user)
}

// Delete 删除用户
func (l *userLogic) Delete(ctx context.Context, id int64) error {
	user, err := l.userDao.GetById(ctx, id)
	if err != nil {
		return err
	}

	opId, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		return err
	}

	if user.Id == opId {
		return errors.New("不能删除自己")
	}

	// 删除用户与角色的关系
	err = l.roleUserDao.DeleteByUserId(ctx, id)
	if err != nil {
		return err
	}

	return l.userDao.Delete(ctx, user.Id)
}

// GetById 根据Id获取用户
func (l *userLogic) GetById(ctx context.Context, id int64) (*systype.UserDataResp, error) {
	user, err := l.userDao.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &systype.UserDataResp{
		Id:          user.Id,
		Username:    user.Username,
		FullName:    user.FullName,
		Avatar:      user.Avatar,
		Email:       user.Email,
		Phone:       user.Phone,
		DeptId:      user.DeptId,
		Status:      user.Status,
		LoginCount:  user.LoginCount,
		LastLoginAt: user.LastLoginAt,
		LastLoginIP: user.LastLoginIP,
		//TenantId:    user.TenantId,
		//OrgId:     user.OrgId,
		Remark:    user.Remark,
		CreatedAt: user.CreatedAt,
		CreatedBy: user.CreatedBy,
		UpdatedAt: user.UpdatedAt,
		UpdatedBy: user.UpdatedBy,
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
			Id:          user.Id,
			Username:    user.Username,
			FullName:    user.FullName,
			Avatar:      user.Avatar,
			Email:       user.Email,
			Phone:       user.Phone,
			DeptId:      user.DeptId,
			Status:      user.Status,
			LoginCount:  user.LoginCount,
			LastLoginAt: user.LastLoginAt,
			LastLoginIP: user.LastLoginIP,
			//TenantId:    user.TenantId,
			//OrgId:     user.OrgId,
			Remark:    user.Remark,
			CreatedAt: user.CreatedAt,
			CreatedBy: user.CreatedBy,
			UpdatedAt: user.UpdatedAt,
			UpdatedBy: user.UpdatedBy,
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

// GetByPhone 根据手机号获取用户
func (l *userLogic) GetByPhone(ctx context.Context, phone string) (*systype.UserDataResp, error) {
	user, err := l.userDao.GetByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}

	if user.Id == 0 {
		err = response.NewError(http.StatusInternalServerError, "用户不存在")
		return nil, err
	}

	return &systype.UserDataResp{
		Id:          user.Id,
		Username:    user.Username,
		Password:    user.Password,
		FullName:    user.FullName,
		Avatar:      user.Avatar,
		Email:       user.Email,
		Phone:       user.Phone,
		DeptId:      user.DeptId,
		Status:      user.Status,
		LoginCount:  user.LoginCount,
		LastLoginAt: user.LastLoginAt,
		LastLoginIP: user.LastLoginIP,
		//TenantId:    user.TenantId,
		//OrgId:     user.OrgId,
		Remark:    user.Remark,
		CreatedAt: user.CreatedAt,
		CreatedBy: user.CreatedBy,
		UpdatedAt: user.UpdatedAt,
		UpdatedBy: user.UpdatedBy,
	}, nil
}

func (l *userLogic) CheckForLogin(ctx context.Context, phone string, password string) (*systype.UserDataResp, error) {
	userInfo, err := l.GetByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}

	// 校验密码是否争取
	if err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password)); err != nil {
		return nil, response.NewError(http.StatusInternalServerError, "密码错误")
	}
	userInfo.Password = ""
	return userInfo, nil
}
