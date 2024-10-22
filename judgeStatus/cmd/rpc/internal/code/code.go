package code

import "oj-micro/common/xcode"

var (
	// RecordNotFound 数据库中不存在该记录
	RecordNotFound = xcode.New(4101, "数据库中不存在该记录")
)
