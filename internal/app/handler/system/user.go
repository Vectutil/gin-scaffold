package system

import (
	syslogic "gin-scaffold/internal/app/logic/system"
	"gin-scaffold/internal/app/response"
	systype "gin-scaffold/internal/app/types/system"
	"gin-scaffold/pkg/mysql"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理器
type UserHandler struct {
}

// NewUserHandler 创建用户Handler实例
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// Create 创建用户
func (h *UserHandler) Create(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       systype.UserCreateReq
		res       = &systype.UserCreateResp{}
		userLogic = syslogic.NewUserLogic(db)
	)

	defer response.HandleDefault(c, res)(&err)

	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}

	// TODO: 从上下文中获取操作者ID
	operatorID := int64(1)

	if err = userLogic.Create(c.Request.Context(), &req, operatorID); err != nil {
		return
	}
}

// Update 更新用户
func (h *UserHandler) Update(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       systype.UserUpdateReq
		res       = &systype.UserUpdateResp{}
		userLogic = syslogic.NewUserLogic(db)
	)

	defer response.HandleDefault(c, res)(&err)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}

	req.ID = id

	// TODO: 从上下文中获取操作者ID
	operatorID := int64(1)

	if err = userLogic.Update(c.Request.Context(), &req, operatorID); err != nil {
		return
	}
}

// Delete 删除用户
func (h *UserHandler) Delete(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		res       = &systype.UserDeleteResp{}
		userLogic = syslogic.NewUserLogic(db)
	)

	defer response.HandleDefault(c, res)(&err)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	// TODO: 从上下文中获取操作者ID
	operatorID := int64(1)

	if err = userLogic.Delete(c.Request.Context(), id, operatorID); err != nil {
		return
	}
}

// GetByID 根据ID获取用户
func (h *UserHandler) GetByID(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		res       = &systype.UserDataResp{}
		userLogic = syslogic.NewUserLogic(db)
	)

	defer response.HandleDefault(c, res)(&err)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	res, err = userLogic.GetByID(c.Request.Context(), id)
}

// List 查询用户列表
func (h *UserHandler) List(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       systype.UserQueryReq
		res       *systype.UserDataListResp
		userLogic = syslogic.NewUserLogic(db)
	)

	defer response.HandleListDefault(c, res)(&err)

	if err = c.ShouldBindQuery(&req); err != nil {
		return
	}

	err = response.ShouldBindForList(c, &req)
	if err != nil {
		return
	}

	res, err = userLogic.GetList(c.Request.Context(), &req)

}
