package pending

import (
	"sai/pkg/util"
	"sai/pkg/workerpool"
)

const (
	goroutineSize = 20
)



var taskPendingHolder =make(map[string]*workerpool.Pool)
// 根据groupid进行创建groutine池
func init() {
	var groupIds = util.GetAllGroupIds()
	for _, groupId := range groupIds {
		pool := workerpool.NewPool(goroutineSize, workerpool.WithBlock(true), workerpool.WithPreAllocWorkers(false))
		taskPendingHolder[groupId] = pool
	}
}

// 根据groupid 获取groutine池
func GetPool(groupId string) *workerpool.Pool {
	return taskPendingHolder[groupId]
}
