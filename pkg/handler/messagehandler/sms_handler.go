package messagehandler

import (
	"context"
	"sai/common"
	"sai/global"
	"sai/model"
	"sai/pkg/util"
	"sai/pkg/util/smsutil"
	"sai/pkg/util/smsutil/tencenSms"
	"time"
)

func init() {
	Handlers[common.CHANNEL_TYPE_SMS] = &SmsHandler{}
}

type SmsHandler struct {
}

func (s *SmsHandler) DoHandler(taskInfo common.TaskInfo) {
	smsParam := common.SmsParam{
		Phones:            taskInfo.Receiver,
		Content:           []string{taskInfo.Content.SmsContent.Content, taskInfo.Content.SmsContent.Expire},
		MessageTemplateId: taskInfo.MessageTemplateId,
		SendAccount:       taskInfo.SendAccount,
	}
	accountOptions := smsutil.GetSmsAccountOptions()
	smsRequest := tencenSms.NewSmsRequest(&accountOptions, tencenSms.WithPhoneNumberSet(smsParam.Phones), tencenSms.WithTemplateParamSet(smsParam.Content))
	smsClient := tencenSms.NewSmsClient(tencenSms.WithRequest(*smsRequest), tencenSms.WithCredential(&accountOptions))
	res, _ := smsClient.Send()
	for _, response := range res.Response.SendStatusSet {
		phone := []byte(*response.PhoneNumber)
		newPhone := phone[len(phone)-11:]
		smsRecord := model.SmsRecord{
			MessageTemplateId: taskInfo.MessageTemplateId,
			Phone:             util.StringToInt64(string(newPhone)),
			SupplierId:        0,
			SupplierName:      "",
			MsgContent:        taskInfo.Content.SmsContent.Content,
			SeriesId:          *response.SerialNo,
			ChargingNum:       *response.Fee,
			ReportContent:     *response.Code,
			Status:            common.SEND_SUCCESS,
			SendDate:          0,
			UpdatedAt:         time.Now(),
			CreatedAt:         time.Now(),
		}
		_ = global.Store.SmsRecord().CreateRecord(context.Background(), &smsRecord)

	}

}
