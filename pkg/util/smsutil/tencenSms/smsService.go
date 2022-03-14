package tencenSms

import (
	tencentCommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
	"sai/common"
	"sai/pkg/logger"
)

type SmsClient struct {
	Credential *tencentCommon.Credential
	Region     string
	Cpf        *profile.ClientProfile
	Request    SmsRequest
}

type Option func(*SmsClient)

func NewSmsClient(options ...func(client *SmsClient)) *SmsClient {
	client := &SmsClient{
		Region: "ap-guangzhou",
		Cpf:    profile.NewClientProfile(),
	}
	for _, option := range options {
		option(client)
	}
	return client

}

func WithRequest(request SmsRequest) Option {
	return func(smsClient *SmsClient) {
		smsClient.Request = request
	}
}

func WithCredential(options *common.SmsAccountOptions) Option {
	return func(smsClient *SmsClient) {
		smsClient.Credential = tencentCommon.NewCredential(options.SecretId, options.SecretKey)
	}
}
func WithCpfReqMethod(method string) Option {
	return func(smsClient *SmsClient) {
		smsClient.Cpf.HttpProfile.ReqMethod = method
	}
}
func WithCpfReqTimeout(timeout int) Option {
	return func(smsClient *SmsClient) {
		smsClient.Cpf.HttpProfile.ReqTimeout = timeout
	}
}
func WithCpfSignMethod(method string) Option {
	return func(smsClient *SmsClient) {
		smsClient.Cpf.SignMethod = method
	}
}

func (s *SmsClient) Send() bool {
	sendClient, _ := sms.NewClient(s.Credential, s.Region, s.Cpf)
	_, err := sendClient.SendSms(s.Request.request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		logger.Warnf("An API error has returned: %s", err)
		return false
	}

	if err != nil {
		logger.Warnf("发送短信失败:%s,requestId:%s", err)
		return false

	}
	logger.Info("发送短信验证码成功")
	return true
}

