package system

import (
	"fmt"
	syslogic "gin-scaffold/internal/app/logic/system"
	sysmodel "gin-scaffold/internal/app/model/system"
	"gin-scaffold/internal/app/response"
	"gin-scaffold/pkg/mysql"
	"gin-scaffold/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MenuHandler 菜单处理器
type MenuHandler struct {
}

// NewMenuHandler 创建菜单Handler实例
func NewMenuHandler() *MenuHandler {
	return &MenuHandler{}
}

// Create 创建菜单
// @title 创建菜单
// @Summary 创建新菜单
// @Description 创建一个新的菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param request body systypes.system.Menu true "菜单创建请求参数"
// @Success 200 {object} systypes.system.Menu "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /menu [post]
func (h *MenuHandler) Create(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       sysmodel.Menu
		res       = &sysmodel.Menu{}
		menuLogic = syslogic.NewMenuLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}

	// 从上下文中获取操作者ID
	operatorID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return
	}

	req.CreatedBy = operatorID
	req.UpdatedBy = operatorID

	if err = menuLogic.Create(c.Request.Context(), &req); err != nil {
		return
	}
}

// Update 更新菜单
// @title 更新菜单
// @Summary 更新菜单信息
// @Description 根据菜单ID更新菜单信息
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "菜单ID"
// @Param request body sysmodel.system.Menu true "菜单更新请求参数"
// @Success 200 {object} sysmodel.system.Menu "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /menu/{id} [put]
func (h *MenuHandler) Update(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       sysmodel.Menu
		res       = &sysmodel.Menu{}
		menuLogic = syslogic.NewMenuLogic(db)
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

	// 从上下文中获取操作者ID
	operatorID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return
	}

	req.UpdatedBy = operatorID

	if err = menuLogic.Update(c.Request.Context(), &req); err != nil {
		return
	}
}

// Delete 删除菜单
// @title 删除菜单
// @Summary 删除指定菜单
// @Description 根据菜单ID删除菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "菜单ID"
// @Success 200 {object} sysmodel.system.Menu "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /menu/{id} [delete]
func (h *MenuHandler) Delete(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		res       = &sysmodel.Menu{}
		menuLogic = syslogic.NewMenuLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	if err = menuLogic.Delete(c.Request.Context(), id); err != nil {
		return
	}
}

// GetByID 根据ID获取菜单
// @title 获取菜单详情
// @Summary 获取指定菜单详情
// @Description 根据菜单ID获取菜单详细信息
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "菜单ID"
// @Success 200 {object} sysmodel.system.Menu "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /menu/{id} [get]
func (h *MenuHandler) GetByID(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		res       = &sysmodel.Menu{}
		menuLogic = syslogic.NewMenuLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	res, err = menuLogic.GetByID(c.Request.Context(), id)
	if err != nil {
		return
	}
}

// List 查询菜单列表
// @title 获取菜单列表
// @Summary 获取菜单列表
// @Description 分页获取菜单列表信息
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} sysmodel.system.Menu "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /menu [get]
func (h *MenuHandler) List(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       interface{}
		res       = &sysmodel.Menu{}
		menuLogic = syslogic.NewMenuLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	menus, total, err := menuLogic.GetList(c.Request.Context(), req)
	if err != nil {
		return
	}
	fmt.Println(menus, total)
	// todo
	// 这里需要根据实际的响应结构体处理 menus 和 total
}

// GetMenuTree 获取菜单树形结构
func (h *MenuHandler) GetMenuTree(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		res       []*sysmodel.Menu
		menuLogic = syslogic.NewMenuLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	res, err = menuLogic.GetMenuTree(c.Request.Context())
	if err != nil {
		return
	}
}
