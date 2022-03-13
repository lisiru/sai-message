package discard

import (

	"sai/common"
	"sai/pkg/util"
)

func IsDiscard(info common.TaskInfo) bool {
	// todo 接入apollo配置中心配置
	configArray:=[]string{"1","2"}

	if _,isExist:=util.SliceContains(configArray, string(info.MessageTemplateId));isExist {
		return true

	}
	return false

}
