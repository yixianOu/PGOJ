package users_profile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oj-micro/users/cmd/api/internal/logic/users_profile"
	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"
)

// RefreshUserRatingHandler 刷新用户判题数据
func RefreshUserRatingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RefreshUserRatingRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := users_profile.NewRefreshUserRatingLogic(r.Context(), svcCtx)
		resp, err := l.RefreshUserRating(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
