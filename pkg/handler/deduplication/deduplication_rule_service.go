package deduplication

import (
	"sai/common"
	"sai/pkg/handler/deduplication/build"
	"sai/pkg/logger"
)

// 平台通用去重逻辑
func Duplication(taskInfo common.TaskInfo)  {
	// todo 从apollo获取去重配置 先写死
	deduplicationConfig:="{\"deduplication_2\":{\"count_num\":1,\"deduplication_time\":300},\"deduplication_1\":{\"count_num\":5}}"
	// 遍历获取当前去重的全部类型
	for _,val:=range common.DeDuplicationTypeList {
		deduplicationParam,err:=build.NewAbstractBuild(val).Build(deduplicationConfig,taskInfo)
		if err!=nil{
			// 记录错误日志
			logger.Infof("构建去重参数失败:%s",err.Error())
			continue

		}
		NewBaseDeDuplication(val).DeDuplication(deduplicationParam)
	}

}
