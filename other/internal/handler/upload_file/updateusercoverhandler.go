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

func UpdateUserCoverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserCoverRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, xcode.InvalidParams.WithDetails(err.Error()))
			return
		}

		file, fileHeader, err := r.FormFile("cover_image")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, xcode.InvalidParams.WithDetails(err.Error()))
			return
		}
		ctx := context.WithValue(r.Context(), "CoverFile", file)
		ctx = context.WithValue(ctx, "CoverFileHeader", fileHeader)

		l := upload_file.NewUpdateUserCoverLogic(ctx, svcCtx)
		resp, err := l.UpdateUserCover(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
