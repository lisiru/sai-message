package tencenSms

import (
	tencentCommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"sai/common"
)

type SmsRequest struct {
	request *sms.SendSmsRequest
}

func NewSmsRequest(options *common.SmsAccountOptions, withOptions ...func(smsRequest *SmsRequest)) *SmsRequest {
	request := sms.NewSendSmsRequest()

	request.SmsSdkAppId = &options.SmsSdkAppId
	request.SignName = &options.SignName
	request.TemplateId = &options.TemplateId
	smsRequest := &SmsRequest{request: request}
	for _, option := range withOptions {
		option(smsRequest)
	}
	return smsRequest

}

type RequestOption func(*SmsRequest)

func WithPhoneNumberSet(phoneSet []string) RequestOption {
	return func(smsRequest *SmsRequest) {
		smsRequest.request.PhoneNumberSet = tencentCommon.StringPtrs(phoneSet)
	}
}

func WithTemplateParamSet(templateSet []string) RequestOption {
	return func(smsRequest *SmsRequest) {
		smsRequest.request.TemplateParamSet = tencentCommon.StringPtrs(templateSet)
	}
}
