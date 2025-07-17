package http_call

type FeishuWebhookRequest struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

type WechatWebhookRequest struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text,omitempty" `
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown,omitempty"`
}
