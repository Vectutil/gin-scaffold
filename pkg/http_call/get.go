package http_call

import (
	"context"
	"io"
	"net/http"
)

// Get 发送一个 HTTP GET 请求并返回响应体
// ctx 上下文，用于控制请求的超时和取消
// url 请求的目标 URL
// headers 请求头
// 返回响应体字节切片和可能出现的错误
func Get(ctx context.Context, url string, headers http.Header) (body []byte, err error) {
	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 创建一个 GET 请求
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	if headers != nil {
		req.Header = headers
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// 确保在函数结束时关闭响应体
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
		}
	}

	// 读取响应体
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
