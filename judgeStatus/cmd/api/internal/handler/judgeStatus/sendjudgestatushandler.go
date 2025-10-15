package judgeStatus

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"oj-micro/common/xcode"
	"oj-micro/judgeStatus/cmd/api/internal/logic/judgeStatus"
	"oj-micro/judgeStatus/cmd/api/internal/svc"
	"oj-micro/judgeStatus/cmd/api/internal/types"
)

func SendJudgeStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendJudgeStatusRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, xcode.InvalidParams.WithDetails(err.Error()))
			return
		}
		l := judgeStatus.NewSendJudgeStatusLogic(r.Context(), svcCtx)
		resp, err := l.SendJudgeStatus(&req)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

//func appendStringsWithSemicolon(stringsToAppend ...string) string {
//	// 使用 strings.Join 方法将多个字符串用 ";" 分隔拼接起来
//	return strings.Join(stringsToAppend, ";")
//}
