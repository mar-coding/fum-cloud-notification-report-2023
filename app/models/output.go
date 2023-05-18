package models

type OutputEmail struct {
	MessageId              int    `json:"id"`
	RequestInsertTimestamp string `json:"inserttime"`
	RequestUpdateTimestamp string `json:"updatetime"`
	Mail_config            int    `json:"mail_config"`
}

type OutputReq struct {
	MessageId           string `json:"message_id"`
	RequestState        int    `json:"state"`
	ItemInsertTimestamp string `json:"inserttime"`
	ItemUpdateTimestamp string `json:"updatetime"`
	Receiver            string `json:"receiver"`
	MailConfigId        int    `json:"mail_config"`
}
