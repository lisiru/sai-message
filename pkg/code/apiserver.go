package code

const  (
	ErrUserNotFound int = iota+110001
	ErrUserPhoneCodeExpire
	ErrUserPhoneCodeMiss
	ErrUserLoginFail
)

const  (

	ErrMessageTemplateNotFound int = iota+110101
)
