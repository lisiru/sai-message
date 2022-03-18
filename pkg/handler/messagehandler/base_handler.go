package messagehandler

import (
	"sai/common"
)

type MessageHandler interface {
	DoHandler(taskInfo common.TaskInfo)
}

type baseHandler struct {
	MessageHandler
}

func NewBaseHandler(channelCode int) *baseHandler {

	return &baseHandler{
		MessageHandler: SelectHandler(channelCode),

	}
}

// 统一的发送接口，各个渠道需要实现这个方法
func (b *baseHandler) Handler(info common.TaskInfo) {
	b.DoHandler(info)
}
