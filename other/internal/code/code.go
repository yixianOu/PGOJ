package code

import "oj-micro/common/xcode"

var (
	// InvalidParams api参数错误
	InvalidParams = xcode.New(03201, "api参数错误")
	// UnauthorizedTokenTimeout jwt授权过期
	UnauthorizedTokenTimeout = xcode.New(03202, "token已过期")
	// UnauthorizedUserNotLogin jwt未授权
	UnauthorizedUserNotLogin = xcode.New(03203, "用户未登录")
	// UnauthorizedUserNotSuperUser 登录用户不是超级用户
	UnauthorizedUserNotSuperUser = xcode.New(03204, "登录用户不是超级用户")
	// FileExists 文件已存在
	FileExists = xcode.New(03205, "文件已存在")
)
