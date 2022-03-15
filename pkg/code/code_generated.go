package code

func init() {
	register(ErrUserNotFound, 404, "该用户不存在")
	register(ErrSendSmsError, 200, "发送短信错误")
	register(ErrUserPhoneCodeExpire,200,"验证码已过期")
	register(ErrUserLoginFail,200,"登录失败，请重试")
	register(ErrParamNotValid,200,"参数错误")
	register(ErrMessageTemplateNotFound,200,"信息模板不存在")

}
