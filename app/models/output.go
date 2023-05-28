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

type PaginationResponseOutputReq struct {
	Meta HATEOASMetaData `json:"metadata"`
	Data []OutputReq     `json:"data"`
}

type PaginationResponseOutputEmail struct {
	Meta HATEOASMetaData `json:"metadata"`
	Data []OutputEmail   `json:"data"`
}

type HATEOASMetaData struct {
	Page       int          `json:"page"`
	PageSize   int          `json:"pageSize"`
	TotalPages int          `json:"totalPages"`
	TotalItems int          `json:"totalItems"`
	Links      HATEOASLinks `json:"_link"`
}

type HATEOASLinks struct {
	Self string `json:"self"`
	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
}
