package users_profile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oj-micro/users/cmd/api/internal/logic/users_profile"
	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"
)

// 列出用户判题信息
func ListProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListProfileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := users_profile.NewListProfileLogic(r.Context(), svcCtx)
		resp, err := l.ListProfile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
