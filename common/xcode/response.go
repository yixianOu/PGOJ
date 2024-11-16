package xcode

import (
	"net/http"
	"oj-micro/common/xcode/types"
)

// ErrHandler 溯源error，得到xcode，返回业务错误码和错误信息
func ErrHandler(err error) (int, any) {
	code := CodeFromError(err)

	return http.StatusOK, types.Status{
		Code:    int32(code.Code()),
		Message: code.Message(),
	}
}
