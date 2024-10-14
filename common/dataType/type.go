package dataType

type RoleLevel int64

const (
	NormalUser RoleLevel = 1
	AdminUser  RoleLevel = 2
)

type CaseCondition string

const (
	Accept              CaseCondition = "ACCEPT"
	WrongAnswer         CaseCondition = "WRONG_ANSWER"
	TimeLimitExceeded   CaseCondition = "TIME_LIMIT_EXCEEDED"
	MemoryLimitExceeded CaseCondition = "MEMORY_LIMIT_EXCEEDED"
	OutputLimitExceeded CaseCondition = "OUTPUT_LIMIT_EXCEEDED"
	RuntimeError        CaseCondition = "RUNTIME_ERROR"
	SegmentationFault   CaseCondition = "SEGMENTATION_FAULT"
	FloatError          CaseCondition = "FLOAT_ERROR"
	UnknownError        CaseCondition = "UNKNOWN_ERROR"
	CompileError        CaseCondition = "COMPILE_ERROR"
)
