package code

import "oj-micro/common/xcode"

var (
	// InvalidParams api参数错误
	InvalidParams = xcode.New(3201, "api参数错误")
	// UnauthorizedTokenTimeout jwt授权过期
	UnauthorizedTokenTimeout = xcode.New(3202, "token已过期")
	// UnauthorizedUserNotLogin jwt未授权
	UnauthorizedUserNotLogin = xcode.New(3203, "用户未登录")
	// UnauthorizedUserNotSuperUser 登录用户不是超级用户
	UnauthorizedUserNotSuperUser = xcode.New(3204, "登录用户不是超级用户")
	// FileExists 文件已存在
	FileExists = xcode.New(3205, "文件已存在")
	// TestcaseExists 测试用例已存在
	TestcaseExists = xcode.New(3206, "测试用例已存在")
)
