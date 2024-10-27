package code

import "oj-micro/common/xcode"

var (
	// UserNameDuplicate 用户名重复
	UserNameDuplicate = xcode.New(1101, "用户名重复")
	// EmailCodeError 验证码错误
	EmailCodeError = xcode.New(1102, "验证码错误")
	// PasswordError 用户密码错误
	PasswordError = xcode.New(1103, "用户密码错误")
	// UserEmailNotExist 该邮箱不存在对应用户
	UserEmailNotExist = xcode.New(1104, "该邮箱不存在对应用户")
	// UserNotFoundError UserNotFound 用户不存在
	UserNotFoundError = xcode.New(1105, "用户不存在")
	// EmailDuplicate 邮箱重复
	EmailDuplicate = xcode.New(1106, "邮箱重复")
)
