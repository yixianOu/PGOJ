package code

import "oj-micro/common/xcode"

var (
	// RecordNotFound 数据库中不存在该记录
	RecordNotFound = xcode.New(2101, "数据库中不存在该记录")
	// TagExist 标签已存在
	TagExist = xcode.New(2102, "标签已存在")

	// TagNameExist 标签名已存在
	TagNameExist = xcode.New(2104, "标签名已存在")
	// ProblemTitleExist 题目标题已存在
	ProblemTitleExist = xcode.New(2105, "题目标题已存在")
	// ProblemCodeExist 题目编号已存在
	ProblemCodeExist = xcode.New(2111, "题目编号已存在")
	// ProblemTagNotExist 该题目不存在你想删除的标签
	ProblemTagNotExist = xcode.New(2106, "该题目不存在你想删除的标签")

	// ProblemSampleExist 该题目已存在你想添加的样例
	ProblemSampleExist = xcode.New(2107, "该题目已存在你想添加的样例")
	// ProblemSampleNotExist 该题目不存在你想删除的样例
	ProblemSampleNotExist = xcode.New(2110, "该题目不存在你想删除的样例")
)
