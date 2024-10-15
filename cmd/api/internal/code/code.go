package code

import "oj-micro/common/xcode"

var (
	// UnauthorizedTokenTimeout JWT过期，需要重新登录
	UnauthorizedTokenTimeout = xcode.New(02201, "token已过期")
	// InvalidParams 参数错误
	InvalidParams = xcode.New(02202, "参数错误")
	// UnauthorizedUserNotSuperUser 无权限，非超级用户
	UnauthorizedUserNotSuperUser = xcode.New(02203, "无权限，非超级用户")
	// UnauthorizedUserNotLogin 无权限，用户未登录
	UnauthorizedUserNotLogin = xcode.New(02204, "无权限，用户未登录")
)
