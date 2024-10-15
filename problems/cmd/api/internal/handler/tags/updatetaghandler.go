package tags

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/api/internal/logic/tags"
	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"
)

func UpdateTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateTagRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, xcode.InvalidParams.WithDetails(err.Error()))
			return
		}

		l := tags.NewUpdateTagLogic(r.Context(), svcCtx)
		resp, err := l.UpdateTag(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
