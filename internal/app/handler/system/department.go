package system

import (
	"gin-scaffold/internal/app/logic/system"
	"gin-scaffold/internal/app/response"
	systype "gin-scaffold/internal/app/types/system"
	"gin-scaffold/internal/middleware/metadata"
	"gin-scaffold/pkg/mysql"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DepartmentHandler 部门处理器
type DepartmentHandler struct {
}

// NewDepartmentHandler 创建部门Handler实例
func NewDepartmentHandler() *DepartmentHandler {
	return &DepartmentHandler{}
}

// Create 创建部门
// @title 创建部门
// @Summary 创建新部门
// @Description 创建一个新的部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param request body systype.DepartmentCreateReq true "部门创建请求参数"
// @Success 200 {object} systype.DepartmentCreateResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /department [post]
func (h *DepartmentHandler) Create(c *gin.Context) {
	var (
		err        error
		db, commit = mysql.GetTrans()
		req        systype.DepartmentCreateReq
		res        = &systype.DepartmentCreateResp{}
		deptLogic  = system.NewDepartmentLogic(db)
	)

	defer func() {
		commit(err)
		response.HandleDefault(c, res)(&err)

	}()

	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}

	// 从上下文中获取操作者Id
	//operatorId, err := utils.GetUserIdFromContext(c)
	//if err != nil {
	//	return
	//}
	operatorId := metadata.GetUserId(c.Request.Context())
	//if err != nil {
	//	return err
	//}

	if err = deptLogic.Create(c.Request.Context(), &req, operatorId); err != nil {
		return
	}
}

// Update 更新部门
// @title 更新部门
// @Summary 更新部门信息
// @Description 根据部门Id更新部门信息
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param id path int true "部门Id"
// @Param request body systype.DepartmentUpdateReq true "部门更新请求参数"
// @Success 200 {object} systype.DepartmentUpdateResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /department/{id} [put]
func (h *DepartmentHandler) Update(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       systype.DepartmentUpdateReq
		res       = &systype.DepartmentUpdateResp{}
		deptLogic = system.NewDepartmentLogic(db)
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

	req.Id = id

	// 从上下文中获取操作者Id
	operatorId := metadata.GetUserId(c.Request.Context())

	if err = deptLogic.Update(c.Request.Context(), &req, operatorId); err != nil {
		return
	}
}

// Delete 删除部门
// @title 删除部门
// @Summary 删除指定部门
// @Description 根据部门Id删除部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param id path int true "部门Id"
// @Success 200 {object} systype.DepartmentDeleteResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /department/{id} [delete]
func (h *DepartmentHandler) Delete(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		res       = &systype.DepartmentDeleteResp{}
		deptLogic = system.NewDepartmentLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	// 从上下文中获取操作者Id
	operatorId := metadata.GetUserId(c.Request.Context())

	if err = deptLogic.Delete(c.Request.Context(), id, operatorId); err != nil {
		return
	}
}

// GetById 根据Id获取部门
// @title 获取部门详情
// @Summary 获取指定部门详情
// @Description 根据部门Id获取部门详细信息
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param id path int true "部门Id"
// @Success 200 {object} systype.DepartmentDataResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /department/{id} [get]
func (h *DepartmentHandler) GetById(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		res       = &systype.DepartmentDataResp{}
		deptLogic = system.NewDepartmentLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	res, err = deptLogic.GetById(c.Request.Context(), id)
	if err != nil {
		return
	}
}

// List 查询部门列表
// @title 获取部门列表
// @Summary 获取部门列表
// @Description 分页获取部门列表信息
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} systype.DepartmentDataListResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /department [get]
func (h *DepartmentHandler) List(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       systype.DepartmentQueryReq
		res       = &systype.DepartmentDataListResp{}
		deptLogic = system.NewDepartmentLogic(db)
	)

	defer func() {
		response.HandleListDefault(c, res)(&err)
	}()

	err = response.ShouldBindForList(c, &req)
	if err != nil {
		return
	}

	res, err = deptLogic.GetList(c.Request.Context(), &req)
	if err != nil {
		return
	}
}

// GetTree 获取部门树
// @title 获取部门树
// @Summary 获取部门树结构
// @Description 获取部门树形结构数据
// @Tags 部门管理
// @Accept json
// @Produce json
// @Success 200 {object} []systype.DepartmentTreeResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /department/tree [get]
func (h *DepartmentHandler) GetTree(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		res       []systype.DepartmentTreeResp
		deptLogic = system.NewDepartmentLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	// 从上下文中获取租户Id
	//tenantId, err := utils.GetTenantIdFromContext(c)
	//if err != nil {
	//	return
	//}

	res, err = deptLogic.GetTree(c.Request.Context())
	if err != nil {
		return
	}
}
