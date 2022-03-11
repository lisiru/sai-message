package pending

import (
	"sai/common"
	"sai/pkg/workerpool"
)

type MessageHandler interface {
	DoHandler(taskInfo common.TaskInfo)
}

func HandlerMessage(info common.TaskInfo) workerpool.Task  {
	return func() {


	}

}