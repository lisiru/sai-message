package common

import "time"

// 定义责任链传输的上下文内容
type ProcessContext struct {
	Code          string // 标识责任链的code
	SendTaskModel SendTaskModel
}

// 发送任务的内容结构
type SendTaskModel struct {
	MessageTemplateId int64          `json:"message_template_id"`
	MessageParamList  []MessageParam `json:"message_param_list"`
	TaskInfo          []TaskInfo     `json:"task_info"`
}

// 请求发送信息的参数
type MessageParam struct {
	Receiver string            `json:"receiver"` // 接收者，多个使用逗号隔开
	Variable map[string]string `json:"variable"` // 消息内容中的可变部分(占位符替换)
	Extra    map[string]string `json:"extra"`    // 扩展参数

}

// 转化成的任务信息结构体
type TaskInfo struct {
	MessageTemplateId int64    `json:"message_template_id"`
	BusinessId        string   `json:"bussiness_id"`
	Receiver          []string `json:"receiver"`
	IdType            int      `json:"id_type"`
	SendChannel       int      `json:"send_channel"`
	MsgType           int      `json:"msg_type"`
	Content           Content  `json:"content"`
	SendAccount       int      `json:"send_account"`
}

// 各种发送渠道内容结构体
type Content struct {
	ImContent              ImContent              `json:"im_content"`
	MiniProgramContent     MiniProgramContent     `json:"mini_program_content"`
	OfficialAccountContent OfficialAccountContent `json:"official_account_content"`
	PushContent            PushContent            `json:"push_content"`
	SmsContent             SmsContent             `json:"sms_content"`
}

type CommonContent struct{}

// api 请求发送的参数结构体
type SendRequestParam struct {
	Code              string       `json:"code"`                                   // 执行业务类型(默认填写 "send")
	MessageTemplateId int64        `json:"message_template_id" binding:"required"` // 消息模板Id
	MessageParam      MessageParam `json:"message_param"`                          // 消息相关的参数

}

// 邮箱发送渠道的发送内容
type EmailContent struct {
	Title   string
	Content string
}

// im 消息发送的内容渠道
type ImContent struct {
}

// 小程序消息渠道的发送内容
type MiniProgramContent struct {
}

// 服务号发送的内容结构体
type OfficialAccountContent struct {
	MapData map[string]string
	Url     string
}

// 推送渠道的内容结构体
type PushContent struct {
}

// 短信渠道发送的内容结构体
type SmsContent struct {
	Content string `json:"content"`
	Expire string `json:"expire"`
	Url     string `json:"url"`
}


// 去重参数结构体
type DeduplicationParam struct {
	TaskInfo          TaskInfo
	DeduplicationTime time.Duration `json:"deduplication_time"`
	CountNum          int           `json:"count_num"`
	// todo 加数据埋点
}

// 发送短信参数结构体
type SmsParam struct {
	Phones            []string
	Content           []string
	MessageTemplateId int64
	SendAccount       int
}

// 短信发送账号options
type SmsAccountOptions struct {
	SecretKey   string `json:"secret-key,omitempty" mapstructure:"secret-key"`
	SecretId    string `json:"secret-id,omitempty" mapstructure:"secret-id"`
	SmsSdkAppId string `json:"sms-sdk-app-id,omitempty" mapstructure:"sms-sdk-app-id"`
	SignName    string `json:"sign-name,omitempty" mapstructure:"sign-name"`
	TemplateId  string `json:"template-id,omitempty" mapstructure:"template-id"`
}
