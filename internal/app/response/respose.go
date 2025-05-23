package response

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`           // 业务码
	Message string      `json:"message"`        // 提示信息
	Data    interface{} `json:"data,omitempty"` // 数据内容（成功时）
}

// Success 成功返回
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// Error 失败返回
func Error(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
	})
}

// HandleDefault ，返回延迟处理函数
func HandleDefault(ctx *gin.Context, res interface{}) func(*error) {
	// 定义延迟处理函数
	handler := func(err *error) {
		if r := recover(); r != nil {
			*err = errors.New(fmt.Sprintf("%v", r))
		}
		if *err != nil {
			Error(ctx, 500, fmt.Sprintf("%v", *err))
			return
		}
		Success(ctx, res)
	}

	return handler
}
