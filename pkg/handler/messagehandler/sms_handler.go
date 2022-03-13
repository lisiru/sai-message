package messagehandler

import "sai/common"

func init() {
	Handlers[common.CHANNEL_TYPE_SMS] = &SmsHandler{}
}

type SmsHandler struct {
}

func (s *SmsHandler) DoHandler(taskInfo common.TaskInfo) {

}
