package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type test struct {
}

func NewTest() *test {
	return &test{}
}

// TestHandle 测试接口
// @title 测试接口Title
// @Summary 测试接口
// @Description 这是一个测试接口，用于演示 Swagger 注释
// @Tags test post
// @Accept json
// @Produce json
// @Param request body Req true "请求参数"
// @Success 200 {object} map[string]interface{} "成功返回"
// @Failure 500 {object} map[string]interface{} "内部错误"
// @Router /test [post]
func (h *test) TestHandle(c *gin.Context) {
	var (
		req = &Req{}
	)
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	time.Sleep(3 * time.Second)
	c.JSON(200, gin.H{
		"message": req.Name,
	})
}

// TestHandleGet 测试接口 get
// @Summary 测试接口 get
// @Description 这是一个测试接口，用于演示 Swagger 注释 get
// @Tags test
// @Accept json
// @Produce json
// @Param name query string false "用户名"
// @Success 200 {object} map[string]interface{} "成功返回"
// @Failure 500 {object} map[string]interface{} "内部错误"
// @Router /test [get]
func (h *test) TestHandleGet(c *gin.Context) {
	var (
		req = &Req{}
	)
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	time.Sleep(3 * time.Second)
	c.JSON(200, gin.H{
		"message": req.Name,
	})
}

// TestHandlePut 测试接口 put
// @Summary 测试接口 put
// @Description 这是一个测试接口，用于演示 Swagger 注释 put
// @Tags test post
// @Accept json
// @Produce json
// @Param request body Req true "请求参数"
// @Success 200 {object} map[string]interface{} "成功返回"
// @Failure 500 {object} map[string]interface{} "内部错误"
// @Router /users/{id} [put]
func (h *test) TestHandlePut(c *gin.Context) {
	var (
		req = &Req{}
	)
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	time.Sleep(3 * time.Second)
	c.JSON(200, gin.H{
		"message": req.Name,
	})
}

// TestHandleDelete 测试接口 delete
// @Summary 测试接口 delete
// @Description 这是一个测试接口，用于演示 Swagger 注释 delete
// @Tags test
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} map[string]interface{} "成功返回"
// @Failure 500 {object} map[string]interface{} "内部错误"
// @Router /users/{id} [delete]
func (h *test) TestHandleDelete(c *gin.Context) {
	// 1. 获取路径参数 id
	id := c.Param("id")
	if id == "" {
		c.JSON(500, gin.H{
			"message": "ID 不能为空",
		})
		return
	}

	// 2. 模拟数据库删除操作
	time.Sleep(3 * time.Second)

	// 3. 返回成功响应
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("用户 %s 删除成功", id),
	})
}

type Req struct {
	Name string `json:"name"`
}
