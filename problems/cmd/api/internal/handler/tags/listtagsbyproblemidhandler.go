package tags

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/api/internal/logic/tags"
	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"
)

func ListTagsByProblemIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListTagsByProblemIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, xcode.InvalidParams.WithDetails(err.Error()))
			return
		}

		l := tags.NewListTagsByProblemIdLogic(r.Context(), svcCtx)
		resp, err := l.ListTagsByProblemId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
