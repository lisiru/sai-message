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

func (b *baseHandler) Handler(info common.TaskInfo) {
	b.DoHandler(info)
}
