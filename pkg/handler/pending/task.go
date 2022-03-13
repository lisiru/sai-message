package pending

import (
	"sai/common"
	"sai/pkg/handler/deduplication"
	"sai/pkg/handler/discard"
	"sai/pkg/handler/messagehandler"
	"sai/pkg/workerpool"
)

func HandlerMessage(info common.TaskInfo) workerpool.Task {
	return func() {
		// 丢弃消息
		if discard.IsDiscard(info) {
			return
		}
		// 对消息进行去重
		deduplication.Duplication(info)
		// 真正进行发送消息
		handler := messagehandler.NewBaseHandler(info.SendChannel)
		handler.Handler(info)

	}

}
