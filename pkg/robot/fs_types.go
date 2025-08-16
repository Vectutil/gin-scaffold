package robot

type FeishuWebhookRequest struct {
	MsgType string                      `json:"msg_type"`
	Content FeishuWebhookRequestContent `json:"content"`
}

type FeishuWebhookRequestContent struct {
	Text string                          `json:"text"`
	Post FeishuWebhookRequestContentPost `json:"post"`
}

type FeishuWebhookRequestContentPost struct {
	ZhCn ZhCn `json:"zh_cn"`
}

type ZhCn struct {
	Title   string          `json:"title"`
	Content [][]ZhCnContent `json:"content"`
}

type ZhCnContent struct {
	Tag      string `json:"tag"`
	Text     string `json:"text,omitempty"`
	Href     string `json:"href,omitempty"`
	UserId   string `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	ImageKey string `json:"image_key,omitempty"`
}
