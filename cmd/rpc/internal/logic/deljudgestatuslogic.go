package logic

import (
	"context"
	"errors"
	"oj-micro/common/xcode"
	"oj-micro/judgeStatus/cmd/rpc/internal/svc"
	"oj-micro/judgeStatus/cmd/rpc/pb"
	"oj-micro/judgeStatus/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelJudgestatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelJudgestatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelJudgestatusLogic {
	return &DelJudgestatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelJudgestatusLogic) DelJudgestatus(in *pb.DelJudgestatusReq) (*pb.DelJudgestatusResp, error) {
	err := l.svcCtx.JudgeStatusModel.Delete(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		l.Logger.Errorf("JudgeStatusModel Delete error: %v", err)
		return nil, xcode.ServerErr
	}
	return &pb.DelJudgestatusResp{}, nil
}
