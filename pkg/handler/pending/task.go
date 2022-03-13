package pending

import (
	"sai/common"
	"sai/pkg/handler/deduplication"
	"sai/pkg/handler/discard"
	"sai/pkg/workerpool"
)

type MessageHandler interface {
	DoHandler(taskInfo common.TaskInfo)
}

func HandlerMessage(info common.TaskInfo) workerpool.Task  {
	return func() {
		// 丢弃消息
		if discard.IsDiscard(info) {
			return
		}
		// 对消息进行去重
		deduplication.Duplication(info)


	}

}


