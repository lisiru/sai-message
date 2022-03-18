package build

import (
	"sai/common"
)

func init() {
	BuildTyepMap[common.DE_DUPLICATION_TYPE_CONTENT] = &ContentDeduplicationBuild{}
}

type ContentDeduplicationBuild struct {
	abstractBuild
}

// 内容去重方式构建参数的方法
func (content *ContentDeduplicationBuild) paramBuild(deduplicationConfig string, info common.TaskInfo) (common.DeduplicationParam, error) {
	deduplicationParam, err := content.getParamsFromConfig(common.DE_DUPLICATION_TYPE_CONTENT, deduplicationConfig, info)
	if err != nil {
		return common.DeduplicationParam{}, err
	}
	// todo 该类型去重下的特定操作
	return deduplicationParam, nil
}
