package smsutil

import (
	"sai/common"
	"sai/global"
)

func GetSmsAccountOptions() common.SmsAccountOptions {
	// 获取配置参数
	// todo 配置改从apollo获取 根据模板中的account获取对应的账号信息
	smsSetting := global.TencenSmsSetting
	accountOptions:=common.SmsAccountOptions{
		SecretKey:   smsSetting.SecretKey,
		SecretId:    smsSetting.SecretId,
		SmsSdkAppId: smsSetting.SmsSdkAppId,
		SignName:    smsSetting.SignName,
		TemplateId:  smsSetting.TemplateId,
	}
	return accountOptions
}
