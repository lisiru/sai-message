package deduplication

import "sai/common"

type FrequencyDeDuplication struct {

}

func init()  {
	TypeMap[common.DE_DUPLICATION_TYPE_FREQUENCY] = &FrequencyDeDuplication{}
}


func (frequency *FrequencyDeDuplication) deduplicationSingleKey(taskInfo common.TaskInfo,receiver string) string{
	return ""


}
