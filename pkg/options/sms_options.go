package options

type SmsOptions struct {
	SecretKey   string `json:"secret-key,omitempty" mapstructure:"secret-key"`
	SecretId    string `json:"secret-id,omitempty" mapstructure:"secret-id"`
	SmsSdkAppId string `json:"sms-sdk-app-id,omitempty" mapstructure:"sms-sdk-app-id"`
	SignName    string `json:"sign-name,omitempty" mapstructure:"sign-name"`
	TemplateId  string `json:"template-id,omitempty" mapstructure:"template-id"`
}

func NewSmsOptions() *SmsOptions {
	return &SmsOptions{
		SecretKey:   "",
		SecretId:    "",
		SmsSdkAppId: "",
		SignName:    "",
		TemplateId:  "",
	}
}
