package robot

type WechatWebhookRequest struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text,omitempty" `
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown,omitempty"`
}
