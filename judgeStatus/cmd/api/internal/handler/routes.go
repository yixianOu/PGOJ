// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	judgeStatus "oj-micro/judgeStatus/cmd/api/internal/handler/judgeStatus"
	"oj-micro/judgeStatus/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 列出判题记录
				Method:  http.MethodGet,
				Path:    "/judge",
				Handler: judgeStatus.ListJudgeStatusHandler(serverCtx),
			},
			{
				// 发送判题请求
				Method:  http.MethodPost,
				Path:    "/judge",
				Handler: judgeStatus.SendJudgeStatusHandler(serverCtx),
			},
			{
				// 删除判题记录
				Method:  http.MethodDelete,
				Path:    "/judge/:judge_id",
				Handler: judgeStatus.DeleteJudgeStatusHandler(serverCtx),
			},
			{
				// 获取判题记录
				Method:  http.MethodGet,
				Path:    "/judge/:judge_id",
				Handler: judgeStatus.GetJudgeStatusHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api4"),
	)
}
