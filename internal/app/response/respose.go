package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin-scaffold/internal/app/types/common"
	"gin-scaffold/pkg/http_call"
	"gin-scaffold/pkg/logger"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

type Response struct {
	Code    int         `json:"code"`           // 业务码
	Message string      `json:"message"`        // 提示信息
	Data    interface{} `json:"data,omitempty"` // 数据内容（成功时）
}

type qwError struct {
	TimeStamp int64       `json:"timeStamp"`
	Code      int         `json:"code"`
	Api       string      `json:"api"`
	Msg       string      `json:"msg"`
	Stack     string      `json:"stack"`
	Request   interface{} `json:"request"`
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
func Error(ctx *gin.Context, err *error, code int, msg interface{}) {

	stack := Stack(*err)
	t := time.Now().UnixNano()

	qerr := qwError{
		TimeStamp: t,
		Code:      code,
		Api:       ctx.Request.URL.Path,
		Msg:       fmt.Sprintf("%+v", msg),
		Stack:     stack,
	}
	body, _ := io.ReadAll(ctx.Request.Body)
	qerr.Request = string(body)
	qerr.Msg = fmt.Sprintf("%+v", *err)

	marshal, _ := json.Marshal(qerr)
	logger.Logger.Error(string(marshal))

	markdown := fmt.Sprintf(`
## 🚨 实时新增接口异常，请相关同事注意 \n
> **时间**：%d  
> **接口**：%s  
> **状态码**：%d  
> **错误信息**：%v

### 📚 堆栈：
%s`,
		t, ctx.Request.URL.Path, code, qerr.Msg, stack)

	http_call.CallQWAssistant(ctx, markdown, http_call.QWRobotMsgTypeMarkdown)

	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: fmt.Sprintf("%v", qerr.Msg),
	})
}

// HandleDefault ，返回延迟处理函数
func HandleDefault(ctx *gin.Context, res interface{}) func(*error, *int) {
	// 定义延迟处理函数
	handler := func(err *error, errCode *int) {
		if r := recover(); r != nil {
			*err = errors.New(fmt.Sprintf("%v", r))
		}
		if *err != nil {
			Error(ctx, err, *errCode, res)
			return
		}
		Success(ctx, res)
	}

	return handler
}
func HandleListDefault(ctx *gin.Context, res common.IBaseListResp) func(*error) {
	// 定义延迟处理函数
	handler := func(err *error) {
		if r := recover(); r != nil {
			*err = errors.New(fmt.Sprintf("%v", r))
		}
		if *err != nil {
			Error(ctx, err, 500, res)
			return
		}
		res.Adjust()
		Success(ctx, res)
	}

	return handler
}

func Stack(err error) string {
	stack := string(debug.Stack())
	// 先替换 \n\t 组合
	all := ">" + strings.ReplaceAll(stack, "\n\t", "\n>")
	return all
}
