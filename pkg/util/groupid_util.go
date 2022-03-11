package util

import "sai/common"

func GetAllGroupIds() []string {
	var groupIds []string
	for _, channel := range common.ChannelTypeEnum {
		for _, messageType := range common.ChannelTypeEnum {
			groupIds = append(groupIds, channel+"."+messageType)
		}

	}
	return groupIds
}

func GetGroupIdByTaskInfo(taskInfo common.TaskInfo) string {
	channelTypeEn := common.ChannelTypeEnum[taskInfo.SendChannel]
	messageTypeEn := common.MessageTypeEnum[taskInfo.MsgType]
	return channelTypeEn + "." + messageTypeEn

}
