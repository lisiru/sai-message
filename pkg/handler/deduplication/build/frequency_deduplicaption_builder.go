package build

import (
	"sai/common"
	"sai/pkg/logger"
)

func init() {
	BuildTyepMap[common.DE_DUPLICATION_TYPE_FREQUENCY] = &FrequencyDeduplicationBuild{}
}

type FrequencyDeduplicationBuild struct {
	abstractBuild
}

func (content *FrequencyDeduplicationBuild) paramBuild(deduplicationConfig string, info common.TaskInfo) (common.DeduplicationParam, error) {
	logger.Info("频率构建参数")
	deduplicationParam, err := content.getParamsFromConfig(common.DE_DUPLICATION_TYPE_FREQUENCY, deduplicationConfig, info)
	if err != nil {
		return common.DeduplicationParam{}, err
	}
	// todo 获取当天剩余的时间
	deduplicationParam.DeduplicationTime = 3000
	return deduplicationParam,nil
	//// todo 该类型去重下的特定操作
	//return deduplicationParam, nil
}
