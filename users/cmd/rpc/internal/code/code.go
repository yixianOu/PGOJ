package code

import "oj-micro/common/xcode"

var (
	// UserNameDuplicate 用户名重复
	UserNameDuplicate = xcode.New(01101, "用户名重复")
	// EmailCodeError 验证码错误
	EmailCodeError = xcode.New(01102, "验证码错误")
	// PasswordError 用户密码错误
	PasswordError = xcode.New(01103, "用户密码错误")
	// UserEmailNotExist 该邮箱不存在对应用户
	UserEmailNotExist = xcode.New(01104, "该邮箱不存在对应用户")
	// UserNotFoundError UserNotFound 用户不存在
	UserNotFoundError = xcode.New(01105, "用户不存在")
)
