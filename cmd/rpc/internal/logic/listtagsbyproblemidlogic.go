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

type ListTagsByProblemIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListTagsByProblemIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTagsByProblemIdLogic {
	return &ListTagsByProblemIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListTagsByProblemIdLogic) ListTagsByProblemId(in *pb.ListTagsByProblemIdReq) (*pb.ListTagsByProblemIdResp, error) {
	builder := l.svcCtx.ProblemTagModel.SelectBuilder()
	result, err := l.svcCtx.ProblemTagModel.FindTagsByProblemId(l.ctx, builder, in.ProblemId)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("find tags by problem id fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}

	tags := make([]*pb.Tag, 0)
	for _, v := range result {
		tag, err := l.svcCtx.TagModel.FindOne(l.ctx, v.TagId)
		if err != nil {
			if errors.Is(err, sqlx.ErrNotFound) {
				return nil, xcode.RecordNotFound
			}
			logx.Errorf("find tag by id fail, err : %v, result : %+v", err, tag)
			return nil, xcode.ServerErr
		}
		tags = append(tags, &pb.Tag{
			TagId:   tag.TagId,
			TagName: tag.TagName,
		})
	}
	return &pb.ListTagsByProblemIdResp{
		Tags: tags,
	}, nil
}
