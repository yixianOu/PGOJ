package judgeStatus

import (
	"context"
	"encoding/json"
	"oj-micro/common/xcode"
	"oj-micro/judgeStatus/cmd/api/internal/svc"
	"oj-micro/judgeStatus/cmd/api/internal/types"
	"oj-micro/judgeStatus/cmd/rpc/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteJudgeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteJudgeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteJudgeStatusLogic {
	return &DeleteJudgeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteJudgeStatusLogic) DeleteJudgeStatus(req *types.DeleteJudgeStatusRequest) (resp *types.DeleteJudgeStatusResponse, err error) {
	//jwt是否过期
	expireTime, err := l.ctx.Value("access_expire").(json.Number).Int64()
	if err != nil || time.Now().Unix() > expireTime {
		return nil, xcode.UnauthorizedTokenTimeout
	}

	result, err := l.svcCtx.JudgeServiceRpc.GetJudgestatusById(l.ctx, &pb.GetJudgestatusByIdReq{
		Id: req.JudgeId,
	})
	if err != nil {
		return nil, err
	}

	userID, err := l.ctx.Value("user_id").(json.Number).Int64()
	if err != nil || userID != result.Judgestatus.UserId {
		return nil, xcode.AccessDenied
	}

	_, err = l.svcCtx.JudgeServiceRpc.DelJudgestatus(l.ctx, &pb.DelJudgestatusReq{
		Id: req.JudgeId,
	})
	if err != nil {
		return nil, err
	}
	return
}
