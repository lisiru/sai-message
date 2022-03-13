package pending

import (
	"sai/pkg/util"
	"sai/pkg/workerpool"
)

const (
	goroutineSize = 20
)



var taskPendingHolder =make(map[string]*workerpool.Pool)

func init() {
	var groupIds = util.GetAllGroupIds()
	for _, groupId := range groupIds {
		pool := workerpool.NewPool(goroutineSize, workerpool.WithBlock(true), workerpool.WithPreAllocWorkers(false))
		taskPendingHolder[groupId] = pool
	}
}

func GetPool(groupId string) *workerpool.Pool {
	return taskPendingHolder[groupId]
}
