package tencenSms

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"sai/pkg/options"
)

type SmsRequest struct {
	request *sms.SendSmsRequest
}

func NewSmsRequest(options *options.SmsOptions, withOptions ...func(smsRequest *SmsRequest)) *SmsRequest {
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
		smsRequest.request.PhoneNumberSet = common.StringPtrs(phoneSet)
	}
}

func WithTemplateParamSet(templateSet []string) RequestOption {
	return func(smsRequest *SmsRequest) {
		smsRequest.request.TemplateParamSet = common.StringPtrs(templateSet)
	}
}
