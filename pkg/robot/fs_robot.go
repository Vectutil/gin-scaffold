package robot

import (
	"bytes"
	"context"
	"encoding/json"
	"gin-scaffold/internal/config"
	"net/http"
)

func SendFeishuRobot(ctx context.Context, data string) {
	// 默认是正式环境
	url := config.Cfg.FSRobot.ErrorRobot
	reqBody := FeishuWebhookRequest{
		MsgType: QWRobotMsgTypeText,
	}
	reqBody.Content.Text = data
	jsonData, _ := json.Marshal(reqBody)
	// 创建 HTTP POST 请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	client.Do(req)
}

func SendFeishuRobotWithUrl(ctx context.Context, url, data, msgType string) {
	// 默认是正式环境
	if msgType == "" {
		msgType = QWRobotMsgTypeText
	}
	reqBody := FeishuWebhookRequest{
		MsgType: msgType,
	}
	switch msgType {
	case MsgTypeText:
		reqBody.Content.Text = data
	case MsgTypePost:
		json.Unmarshal([]byte(data), &reqBody.Content.Post.ZhCn)
	}
	jsonData, _ := json.Marshal(reqBody)
	// 创建 HTTP POST 请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	client.Do(req)
	//do, err := client.Do(req)
	//if err != nil {
	//	logger.Logger.Error(err.Error())
	//}
	//defer do.Body.Close()
	//// 读取响应体
	//body, err := io.ReadAll(do.Body)
	//fmt.Println(string(body))
}
