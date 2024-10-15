package upload_file

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"oj-micro/common/xcode"
	"oj-micro/other/internal/logic/upload_file"
	"oj-micro/other/internal/svc"
	"oj-micro/other/internal/types"
)

func AddTestCaseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddTestCaseRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, xcode.InvalidParams.WithDetails(err.Error()))
			return
		}

		file, fileHeader, err := r.FormFile("sample_output")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, xcode.InvalidParams.WithDetails(err.Error()))
			return
		}
		ctx := context.WithValue(r.Context(), "outputFile", file)
		ctx = context.WithValue(ctx, "outputFileHeader", fileHeader)

		file, fileHeader, err = r.FormFile("sample_input")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, xcode.InvalidParams.WithDetails(err.Error()))
			return
		}
		ctx = context.WithValue(ctx, "inputFile", file)
		ctx = context.WithValue(ctx, "inputFileHeader", fileHeader)

		l := upload_file.NewAddTestCaseLogic(ctx, svcCtx)
		resp, err := l.AddTestCase(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
