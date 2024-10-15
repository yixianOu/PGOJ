package other

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/api/internal/logic/other"
	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"
)

func SendEmailToUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendEmailToUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, xcode.InvalidParams.WithDetails(err.Error()))
			return
		}

		l := other.NewSendEmailToUserLogic(r.Context(), svcCtx)
		resp, err := l.SendEmailToUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
