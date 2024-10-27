package xcode

var (
	// OK 成功
	OK = add(0, "OK")
	// RequestErr 请求参数错误
	RequestErr = add(400, "INVALID_ARGUMENT")
	// Unauthorized 未认证
	Unauthorized = add(401, "UNAUTHENTICATED")
	// AccessDenied 不是数据属主
	AccessDenied = add(403, "PERMISSION_DENIED")
	// NotFound 未找到记录
	NotFound = add(404, "NOT_FOUND")
	// MethodNotAllowed 请求方法不允许
	MethodNotAllowed = add(405, "METHOD_NOT_ALLOWED")
	// Canceled 请求取消
	Canceled = add(498, "CANCELED")
	// ServerErr 服务端错误
	ServerErr = add(500, "INTERNAL_ERROR")
	// ServiceUnavailable 服务不可用
	ServiceUnavailable = add(503, "UNAVAILABLE")
	// Deadline 服务器处理超时
	Deadline = add(504, "DEADLINE_EXCEEDED")
	// LimitExceed 资源耗尽
	LimitExceed = add(509, "RESOURCE_EXHAUSTED")
	// UnauthorizedTokenTimeout 未认证token超时
	UnauthorizedTokenTimeout = add(600, "UNAUTHORIZED_TOKEN_TIMEOUT ")
	// UnauthorizedUserNotLogin 未认证用户未登录
	UnauthorizedUserNotLogin = add(601, "NOT_LOGIN")
	// InvalidParams 参数错误
	InvalidParams = add(602, "INVALID_PARAMS")
	// UnauthorizedUserNotSuperUser 未认证用户不是超级用户
	UnauthorizedUserNotSuperUser = add(603, "UNAUTHORIZED")
	// RecordNotFound 记录未找到
	RecordNotFound = add(701, "RECORD_NOT_FOUND")
	// RequestTimeout 请求超时
	RequestTimeout = add(702, "REQUEST_TIMEOUT")
	//NoLogin                      = add(101, "NOT_LOGIN")
)
