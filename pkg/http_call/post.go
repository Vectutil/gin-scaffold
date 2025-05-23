package http_call

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	ContentTypeForm = "application/x-www-form-urlencoded"
	ContentTypeJSON = "application/json"
)

// 发起 POST 请求 (支持 form 和 json)
func HttpPost(urlStr string, data interface{}, contentType string, headers map[string]string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	var body io.Reader
	switch contentType {
	case ContentTypeJSON:
		jsonData, err := json.Marshal(data)
		if err != nil {
			return "", err
		}
		body = bytes.NewBuffer(jsonData)
	case ContentTypeForm:
		formData, ok := data.(map[string]string)
		if !ok {
			return "", fmt.Errorf("form data 格式错误，应为 map[string]string")
		}
		values := url.Values{}
		for k, v := range formData {
			values.Set(k, v)
		}
		body = strings.NewReader(values.Encode())
	default:
		return "", fmt.Errorf("不支持的 Content-Type: %s", contentType)
	}

	req, err := http.NewRequest("POST", urlStr, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", contentType)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	return string(respBody), err
}
