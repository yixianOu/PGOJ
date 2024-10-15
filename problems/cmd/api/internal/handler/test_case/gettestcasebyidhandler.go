package test_case

import (
	"net/http"
	"oj-micro/common/xcode"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oj-micro/problems/cmd/api/internal/logic/test_case"
	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"
)

func GetTestCaseByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTestCaseByIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, xcode.InvalidParams.WithDetails(err.Error()))
			return
		}

		l := test_case.NewGetTestCaseByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetTestCaseById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
