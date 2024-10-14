package code

import "oj-micro/common/xcode"

var (
	// Unauthorized 鉴权失败
	Unauthorized = xcode.New(01204, "鉴权失败")
	// EmailFormatError 邮箱格式错误
	EmailFormatError = xcode.New(01205, "邮箱格式错误")
	// SendEmailError 发送邮件失败
	SendEmailError = xcode.New(01206, "发送邮件失败")
	// RedisSaveCodeError redis保存验证码失败
	RedisSaveCodeError = xcode.New(01207, "redis保存验证码失败")

	// InvalidParams 参数错误
	InvalidParams = xcode.New(01210, "参数错误")
)
