package users_info

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/api/internal/logic/users_info"
	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"
)

func ListUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, xcode.InvalidParams.WithDetails(err.Error()))
			return
		}

		l := users_info.NewListUserLogic(r.Context(), svcCtx)
		resp, err := l.ListUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
