package common

type ProcessContext struct {
	Code          string // 标识责任链的code
	SendTaskModel SendTaskModel
}

type SendTaskModel struct {
	MessageTemplateId int64          `json:"message_template_id"`
	MessageParamList  []MessageParam `json:"message_param_list"`
	TaskInfo          []TaskInfo     `json:"task_info"`
}
type MessageParam struct {
	Receiver string            `json:"receiver"` // 接收者，多个使用逗号隔开
	Variable map[string]string `json:"variable"` // 消息内容中的可变部分(占位符替换)
	Extra    map[string]string `json:"extra"`    // 扩展参数

}

type TaskInfo struct {
	MessageTemplateId int64       `json:"message_template_id"`
	BusinessId        string      `json:"bussiness_id"`
	Receiver          []string    `json:"receiver"`
	IdType            int         `json:"id_type"`
	SendChannel       int         `json:"send_channel"`
	MsgType           int         `json:"msg_type"`
	Content           interface{} `json:"content"`
	SendAccount       int         `json:"send_account"`
}

type SendRequestParam struct {
	Code              string       `json:"code"`                                   // 执行业务类型(默认填写 "send")
	MessageTemplateId int64        `json:"message_template_id" binding:"required"` // 消息模板Id
	MessageParam      MessageParam `json:"message_param"`                          // 消息相关的参数

}

type EmailContent struct {
	Title   string
	Content string
}

type ImContent struct {
}

type MiniProgramContent struct {
}

type OfficialAccountContent struct {
	MapData map[string]string
	Url     string
}

type PushContent struct {
}

type SmsContent struct {
	Content string
	Url     string
}

var ChannelContentMap = map[int]interface{} {
	CHANNEL_TYPE_IM: ImContent{},
	CHANNEL_TYPE_EMAIL: EmailContent{},
	CHANNEL_OFFICIAL_ACCOUNT: OfficialAccountContent{},
	CHANNEL_TYPE_PUSH: PushContent{},
	CHANNEL_TYPE_SMS: SmsContent{},
}
