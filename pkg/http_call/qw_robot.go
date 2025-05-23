package http_call

import (
	"bytes"
	"context"
	"encoding/json"
	"gin-scaffold/internal/config"
	"net/http"
)

func CallQWAssistant(ctx context.Context, data string) {

	// 默认是正式环境
	url := config.Cfg.WXRobot.ErrorRobot
	reqBody := WechatWebhookRequest{
		MsgType: "text",
	}

	reqBody.Text.Content = data
	jsonData, _ := json.Marshal(reqBody)
	// 创建 HTTP POST 请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	client.Do(req)
}

type WechatWebhookRequest struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}
