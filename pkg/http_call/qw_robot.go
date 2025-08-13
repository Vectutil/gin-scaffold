package http_call

import (
	"bytes"
	"context"
	"encoding/json"
	"gin-scaffold/internal/config"
	"net/http"
)

const (
	QWRobotMsgTypeText     = "text"
	QWRobotMsgTypeMarkdown = "markdown"
)

func CallQWAssistant(ctx context.Context, data, QWRobotMsgType string) {
	SendFeishuRobot(ctx, data)
	//SendQWRobot(ctx, data, QWRobotMsgType)
}

func SendFeishuRobot(ctx context.Context, data string) {
	// 默认是正式环境
	url := config.Cfg.FSRobot.ErrorRobot
	reqBody := FeishuWebhookRequest{
		MsgType: QWRobotMsgTypeText,
		Content: struct {
			Text string `json:"text"`
		}{
			Text: data,
		},
	}
	jsonData, _ := json.Marshal(reqBody)
	// 创建 HTTP POST 请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	client.Do(req)
}

func SendFeishuRobotWithUrl(ctx context.Context, url, data string) {
	// 默认是正式环境
	reqBody := FeishuWebhookRequest{
		MsgType: QWRobotMsgTypeText,
		Content: struct {
			Text string `json:"text"`
		}{
			Text: data,
		},
	}
	jsonData, _ := json.Marshal(reqBody)
	// 创建 HTTP POST 请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	client.Do(req)
}

func SendQWRobot(ctx context.Context, data, QWRobotMsgType string) {
	// 默认是正式环境
	url := config.Cfg.WXRobot.ErrorRobot
	reqBody := WechatWebhookRequest{
		MsgType: QWRobotMsgType,
	}

	switch QWRobotMsgType {
	case QWRobotMsgTypeText:
		reqBody.Text.Content = data
	case QWRobotMsgTypeMarkdown:
		reqBody.Markdown.Content = data
	}

	jsonData, _ := json.Marshal(reqBody)
	// 创建 HTTP POST 请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	client.Do(req)
}
