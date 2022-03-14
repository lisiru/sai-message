package messagehandler

import (
	"sai/common"
	"sai/pkg/util/smsutil"
	"sai/pkg/util/smsutil/tencenSms"
)

func init() {
	Handlers[common.CHANNEL_TYPE_SMS] = &SmsHandler{}
}

type SmsHandler struct {
}

func (s *SmsHandler) DoHandler(taskInfo common.TaskInfo) {
	smsParam:=common.SmsParam{
		Phones:            taskInfo.Receiver,
		Content:           []string{taskInfo.Content.SmsContent.Content,taskInfo.Content.SmsContent.Expire},
		MessageTemplateId: taskInfo.MessageTemplateId,
		SendAccount:       taskInfo.SendAccount,
	}
	accountOptions:=smsutil.GetSmsAccountOptions()
	smsRequest:=tencenSms.NewSmsRequest(&accountOptions,tencenSms.WithPhoneNumberSet(smsParam.Phones),tencenSms.WithTemplateParamSet(smsParam.Content))
	smsClient:=tencenSms.NewSmsClient(tencenSms.WithRequest(*smsRequest),tencenSms.WithCredential(&accountOptions))
	smsClient.Send()


}
