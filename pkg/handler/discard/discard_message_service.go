package discard

import (

	"sai/common"
	"sai/pkg/logger"
	"sai/pkg/util"
)

func IsDiscard(info common.TaskInfo) bool {
	// todo 接入apollo配置中心配置
	configArray:=[]string{"1","2"}
	logger.Infof("info messageTemplateId:%d",info.MessageTemplateId)

	if _,isExist:=util.SliceContains(configArray, string(info.MessageTemplateId));isExist {

		return true

	}
	logger.Info("跳过")
	return false

}
