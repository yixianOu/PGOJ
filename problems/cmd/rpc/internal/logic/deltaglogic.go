package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTagLogic {
	return &DelTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTagLogic) DelTag(in *pb.DelTagReq) (*pb.DelTagResp, error) {
	builder := l.svcCtx.ProblemTagModel.SelectBuilder()
	result, err := l.svcCtx.ProblemTagModel.FindProblemsByTagId(l.ctx, builder, in.Id, 0, 0)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("find problems by tag id fail, err : %v", err)
		return nil, xcode.ServerErr
	}

	for _, v := range result {
		err = l.svcCtx.ProblemTagModel.Delete(l.ctx, v.Id)
		if err != nil {
			if errors.Is(err, sqlx.ErrNotFound) {
				return nil, xcode.RecordNotFound
			}
			logx.Errorf("delete problem tag fail, err : %v", err)
			return nil, xcode.ServerErr
		}
	}

	err = l.svcCtx.TagModel.Delete(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("delete tag fail, err : %v", err)
		return nil, xcode.ServerErr
	}
	return &pb.DelTagResp{}, nil
}
