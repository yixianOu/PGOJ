package logic

import (
	"context"
	"errors"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"
	"oj-micro/problems/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProblemLogic {
	return &DelProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelProblemLogic) DelProblem(in *pb.DelProblemReq) (*pb.DelProblemResp, error) {
	//查找问题标签
	builder := l.svcCtx.ProblemTagModel.SelectBuilder()
	result, err := l.svcCtx.ProblemTagModel.FindTagsByProblemId(l.ctx, builder, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		l.Logger.Errorf("find tags by problem id fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}
	//删除问题标签
	for _, tag := range result {
		err = l.svcCtx.ProblemTagModel.Delete(l.ctx, tag.Id)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return nil, xcode.RecordNotFound
			}
			l.Logger.Errorf("delete problem tag fail, err : %v, result : %+v", err, tag)
			return nil, xcode.ServerErr
		}
	}
	//查找
	problemData, err := l.svcCtx.ProblemdataModel.FindOneByProblemId(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		l.Logger.Errorf("find problem data by problem id fail, err : %v, result : %+v", err, problemData)
		return nil, xcode.ServerErr
	}
	//删除问题数据
	err = l.svcCtx.ProblemdataModel.Delete(l.ctx, problemData.ProblemdataId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		l.Logger.Errorf("delete problem data fail, err : %v, result : %+v", err, problemData)
		return nil, xcode.ServerErr
	}
	//删除问题
	err = l.svcCtx.ProblemModel.Delete(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		l.Logger.Errorf("delete problem fail, err : %v, result : %+v", err, in.Id)
		return nil, xcode.ServerErr
	}

	return &pb.DelProblemResp{}, nil
}
