package models

type WebHookReqBody struct {
	Message struct {
		Message_id int `json:message_id`
		From       struct {
			Username string `json:"username"`
			Id       int    `json:"Id"`
		} `json:"from"`
		Text string `json:"text"`
		Chat struct {
			ID int `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type SendMessageReqBody struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}
