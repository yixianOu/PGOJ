package judgeStatus

import (
	"net/http"
	"oj-micro/common/xcode"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oj-micro/judgeStatus/cmd/api/internal/logic/judgeStatus"
	"oj-micro/judgeStatus/cmd/api/internal/svc"
	"oj-micro/judgeStatus/cmd/api/internal/types"
)

func ListJudgeStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListJudgeStatusRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, xcode.InvalidParams.WithDetails(err.Error()))
			return
		}

		l := judgeStatus.NewListJudgeStatusLogic(r.Context(), svcCtx)
		resp, err := l.ListJudgeStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
