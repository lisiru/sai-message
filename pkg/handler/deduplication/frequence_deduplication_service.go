package deduplication

import (
	"sai/common"
	"sai/pkg/logger"
)

type FrequencyDeDuplication struct {

}

func init()  {
	TypeMap[common.DE_DUPLICATION_TYPE_FREQUENCY] = &FrequencyDeDuplication{}
}

/**
频率去重方式构建key 的方法
 */
func (frequency *FrequencyDeDuplication) deduplicationSingleKey(taskInfo common.TaskInfo,receiver string) string{
	logger.Info("frequency")
	return ""


}
