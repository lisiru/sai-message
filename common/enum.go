package common

const  (
	ID_TYPE_USER_ID int= iota+1
	ID_TYPE_DID
	ID_TYPE_PHONE
	ID_TYPE_OPEN_ID
	ID_TYPE_EMAIL

)


// 发送渠道枚举
const  (
	CHANNEL_TYPE_IM int  = iota+1
	CHANNEL_TYPE_PUSH
	CHANNEL_TYPE_SMS
	CHANNEL_TYPE_EMAIL
	CHANNEL_OFFICIAL_ACCOUNT
	MINI_PROGRAM
)

// 消息类型枚举
const  (
	MESSAGE_TYPE_NOTICE int = iota+1 // 通知类
	MESSAGE_TYPE_MARKETING // 营销类
	MESSAGE_TYPE_AUTH_CODE // 验证码

)

// 消息类型映射枚举
var  MessageTypeEnum = map[int]string{
	MESSAGE_TYPE_NOTICE: "notice",
	MESSAGE_TYPE_MARKETING: "marketing",
	MESSAGE_TYPE_AUTH_CODE: "auth_code",
}

// 发送渠道映射枚举
var ChannelTypeEnum = map[int]string{
	CHANNEL_TYPE_IM: "im",
	CHANNEL_TYPE_SMS: "sms",
	CHANNEL_TYPE_PUSH: "push",
	CHANNEL_OFFICIAL_ACCOUNT: "official_account",
	CHANNEL_TYPE_EMAIL: "email",
	MINI_PROGRAM: "mini_program",

}

const  (
	TRUE=1
	FALSE=0
)

// 去重类型枚举
const  (
	DE_DUPLICATION_TYPE_CONTENT= iota+1
	DE_DUPLICATION_TYPE_FREQUENCY
)

var DeDuplicationTypeList = []int{DE_DUPLICATION_TYPE_CONTENT,DE_DUPLICATION_TYPE_FREQUENCY}

// 短信状态信息

const  (
	SEND_SUCCESS int8= iota+1 //调用渠道接口发送成功
	RECEIVE_SUCCESS  //用户收到短信(收到渠道短信回执，状态成功)
	RECEIVE_FAIL //用户收不到短信(收到渠道短信回执，状态失败)
)









