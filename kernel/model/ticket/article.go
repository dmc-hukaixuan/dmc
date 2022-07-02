package ticket

type ArticleGet struct {
	TicketID               int64  `json:"ticketID"`
	ArticleID              int64  `json:"articleID"` // Optional filters, these can be combined
	CommunicationChannel   string `json:"communicationChannel"`
	CommunicationChannelID int    `json:"communicationChannelID"`
	SenderType             string `json:"senderType"`
	SenderTypeID           string `json:"senderTypeID"`
	IsVisibleForCustomer   string `json:"isVisibleForCustomer"`
	OnlyFirst              int    `json:"onlyFirst"` // After filtering, you can also limit to first or last found article only:
	OnlyLast               string `json:"onlyLast"`
}

type ArticleDataMimeCreate struct {
	ID                              int64  `json:"id"`
	TicketID                        int64  `json:"ticketID" `
	SenderTypeID                    int    `json:"articleID"` // Optional filters, these can be combined
	SenderType                      string `json:"SenderType"`
	IsVisibleForCustomer            int    `json:"isVisibleForCustomer"`
	UserID                          int    `json:"senderType"`
	From                            string `json:"senderTypeID"`
	To                              string `json:"to"`
	Cc                              string `json:"cc"` // After filtering, you can also limit to first or last found article only:
	Bcc                             string `json:"bcc"`
	ReplyTo                         string `json:"ReplyTo"`
	Subject                         string `json:"Subject"`
	Body                            string `json:"Body"`
	MimeType                        string `json:"mimeType"`
	Charset                         string `json:"charset"`
	CommunicationChannel            string `json:"communicationChannel"`
	MessageID                       int64  `json:"MessageID"`
	InReplyTo                       string `json:"InReplyTo"`
	References                      string `json:"References"`
	ContentType                     string `json:"ContentType"`
	HistoryType                     string `json:"HistoryType"`
	HistoryComment                  string `json:"HistoryComment"`
	NoAgentNotify                   string `json:"NoAgentNotify"`
	AutoResponseType                string `json:"AutoResponseType"`
	ForceNotificationToUserID       string `json:"ForceNotificationToUserID"`
	ExcludeNotificationToUserID     string `json:"ExcludeNotificationToUserID"`
	ExcludeMuteNotificationToUserID string `json:"ExcludeMuteNotificationToUserID"`
}

type ArticleDataMime struct {
	ID           int64  `json:"id" gorm:"id"`
	ArticleID    int64  `json:"article_id" gorm:"article_id"`
	TicketID     int64  `json:"ticket_id" gorm:"ticket_id"`
	From         string `json:"a_from" gorm:"a_from"` // Optional filters, these can be combined
	ReplyTo      string `json:"a_reply_to" gorm:"a_reply_to"`
	To           string `json:"a_to" gorm:"a_to"`
	Cc           string `json:"a_cc" gorm:"a_cc"`
	Bcc          string `json:"a_bcc" gorm:"a_bcc"`
	Subject      string `json:"a_subject" gorm:"a_subject"`
	MessageID    int    `json:"a_message_id" gorm:"a_message_id"` // After filtering, you can also limit to first or last found article only:
	MessageIDMD5 string `json:"a_message_id_md5" gorm:"a_message_id_md5"`
	InReplyTo    string `json:"a_in_reply_to" gorm:"a_in_reply_to"`
	References   string `json:"a_references" gorm:"a_references"`
	ContentType  string `json:"a_content_type" gorm:"a_content_type"`
	ContentSize  int    `json:"a_content_size" gorm:"content_size"`
	Body         string `json:"a_body" gorm:"content"`
	IncomingTime int64  `json:"incoming_time" gorm:"incoming_time"`
	ContentPath  string `json:"content_path" gorm:"content_path"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;"`
	CreateBy     int    `json:"createBy" gorm:"column:create_by;"`
	CreateByName string `gorm:"<-:false" json:"create_by_name"`
	ChangeTime   string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy     int    `json:"changeBy" gorm:"column:change_by;"`
	ChangeByName string `gorm:"<-:false" json:"change_by_name,omitempty"`
}

type ArticleRichText struct {
	ID           int64  `json:"id" gorm:"id"`
	ArticleID    int64  `json:"article_id" gorm:"article_id"`
	TicketID     int64  `json:"ticket_id" gorm:"ticket_id"`
	ContentType  string `json:"a_content_type" gorm:"a_content_type"`
	ContentSize  int    `json:"a_content_size" gorm:"content_size"`
	Body         string `json:"a_body" gorm:"content"`
	ContentPath  string `json:"content_path" gorm:"content_path"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;"`
	CreateBy     int    `json:"createBy" gorm:"column:create_by;"`
	CreateByName string `gorm:"<-:false" json:"create_by_name"`
	ChangeTime   string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy     int    `json:"changeBy" gorm:"column:change_by;"`
	ChangeByName string `gorm:"<-:false" json:"change_by_name,omitempty"`
}

type ArticleData struct {
	ID         int64  `json:"id"`
	TicketID   int64  `json:"ticketID" `
	SenderType string `json:"senderType"`
}

type MetaArticle struct {
	ID                      int64  `json:"id" gorm:"id"`
	TicketID                int64  `json:"ticketID" gorm:"column:ticket_id;"`
	SenderType              int    `json:"senderType" gorm:"column:article_sender_type_id;"`
	IsVisibleForCustomer    int    `json:"isVisibleForCustomer" gorm:"column:is_visible_for_customer;"`
	CommunicationChannelID  int    `json:"communicationChannel" gorm:"column:communication_channel_id;"`
	SearchIndexNeedsRebuild int    `json:"searchIndexNeedsRebuild" gorm:"column:search_index_needs_rebuild;"`
	InsertFingerprint       string `json:"insertFingerprint" gorm:"column:insert_fingerprint;"`
	UserID                  int    `json:"UserID"`
	CreateTime              string `json:"createTime" gorm:"column:create_time;"`
	CreateBy                int    `json:"createBy" gorm:"column:create_by;"`
	CreateByName            string `gorm:"<-:false" json:"create_by_name"`
	ChangeTime              string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy                int    `json:"changeBy" gorm:"column:change_by;"`
	ChangeByName            string `gorm:"<-:false" json:"change_by_name,omitempty"`
}

type ArticleDataMimeAttachment struct {
	ID                 int64  `json:"id" gorm:"id"`
	TicketID           int64  `json:"ticketID" gorm:"column:ticket_id;"`
	ArticleID          int64  `json:"article_id" gorm:"column:article_id;"`
	Filename           string `json:"filename" gorm:"column:filename;"`
	ContentSize        int64  `json:"content_size" gorm:"column:content_size;"`
	ContentType        string `json:"content_type" gorm:"column:content_type;"`
	ContentID          string `json:"content_id" gorm:"column:content_id;"`
	ContentAlternative string `json:"content_alternative" gorm:"column:content_alternative;"`
	Disposition        string `json:"disposition" gorm:"column:disposition;"`
	Content            string `json:"content" gorm:"column:content;"`
	CreateTime         string `json:"createTime" gorm:"column:create_time;"`
	CreateBy           int    `json:"createBy" gorm:"column:create_by;"`
	CreateByName       string `gorm:"<-:false" json:"create_by_name"`
	ChangeTime         string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy           int    `json:"changeBy" gorm:"column:change_by;"`
	ChangeByName       string `gorm:"<-:false" json:"change_by_name,omitempty"`
}
