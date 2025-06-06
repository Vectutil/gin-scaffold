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
// @title 创建用户
// @Summary 创建新用户
// @Description 创建一个新的系统用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body systype.UserCreateReq true "用户创建请求参数"
// @Success 200 {object} systype.UserCreateResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /user [post]
func (h *UserHandler) Create(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       systype.UserCreateReq
		res       = &systype.UserCreateResp{}
		userLogic = syslogic.NewUserLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

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
// @title 更新用户
// @Summary 更新用户信息
// @Description 根据用户ID更新用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param request body systype.UserUpdateReq true "用户更新请求参数"
// @Success 200 {object} systype.UserUpdateResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /user/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       systype.UserUpdateReq
		res       = &systype.UserUpdateResp{}
		userLogic = syslogic.NewUserLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

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
// @title 删除用户
// @Summary 删除指定用户
// @Description 根据用户ID删除用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} systype.UserDeleteResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /user/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		res       = &systype.UserDeleteResp{}
		userLogic = syslogic.NewUserLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

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
// @title 获取用户详情
// @Summary 获取指定用户详情
// @Description 根据用户ID获取用户详细信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} systype.UserDataResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /user/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		res       = &systype.UserDataResp{}
		userLogic = syslogic.NewUserLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	res, err = userLogic.GetByID(c.Request.Context(), id)
}

// List 查询用户列表
// @title 获取用户列表
// @Summary 获取用户列表
// @Description 分页获取用户列表信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} systype.UserDataListResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /user [get]
func (h *UserHandler) List(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       systype.UserQueryReq
		res       = &systype.UserDataListResp{}
		userLogic = syslogic.NewUserLogic(db)
	)

	defer func() {
		response.HandleListDefault(c, res)(&err)
	}()

	if err = c.ShouldBindQuery(&req); err != nil {
		return
	}

	err = response.ShouldBindForList(c, &req)
	if err != nil {
		return
	}

	res, err = userLogic.GetList(c.Request.Context(), &req)
	if err != nil {
		return
	}

}
