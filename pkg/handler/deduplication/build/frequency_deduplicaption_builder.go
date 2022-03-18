package build

import (
	"sai/common"
	"sai/pkg/logger"
	"time"
)

func init() {
	BuildTyepMap[common.DE_DUPLICATION_TYPE_FREQUENCY] = &FrequencyDeduplicationBuild{}
}

type FrequencyDeduplicationBuild struct {
	abstractBuild
}

// 频率去重方式构建参数的方法
func (content *FrequencyDeduplicationBuild) paramBuild(deduplicationConfig string, info common.TaskInfo) (common.DeduplicationParam, error) {
	logger.Info("频率构建参数")
	deduplicationParam, err := content.getParamsFromConfig(common.DE_DUPLICATION_TYPE_FREQUENCY, deduplicationConfig, info)
	if err != nil {
		return common.DeduplicationParam{}, err
	}
	// todo 获取当天剩余的时间
	deduplicationParam.DeduplicationTime = 3000*time.Second
	return deduplicationParam,nil
	//// todo 该类型去重下的特定操作
	//return deduplicationParam, nil
}
