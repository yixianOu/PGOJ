package logic

import (
	"context"
	"errors"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/internal/code"
	"oj-micro/problems/model"

	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProblemLogic {
	return &UpdateProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProblemLogic) UpdateProblem(in *pb.UpdateProblemReq) (*pb.UpdateProblemResp, error) {
	if in.TagIds != nil {
		if in.TagOperation {
			for _, v := range in.TagIds {
				result, err := l.svcCtx.ProblemTagModel.Insert(l.ctx, &model.ProblemTag{TagId: v, ProblemId: in.ProblemId})
				if err != nil {
					l.Logger.Errorf("insert problem tag fail, err : %v, result : %+v", err, result)
				}
			}
		} else {
			for _, v := range in.TagIds {
				result, err := l.svcCtx.ProblemTagModel.FindOneByProblemIdTagId(l.ctx, in.ProblemId, v)
				if err != nil {
					if errors.Is(err, model.ErrNotFound) {
						return nil, code.ProblemTagNotExist
					}
					l.Logger.Errorf("find problem tag fail, err : %v, result : %+v", err, result)
					return nil, xcode.ServerErr
				}
				err = l.svcCtx.ProblemTagModel.Delete(l.ctx, result.Id)

				if err != nil {
					if errors.Is(err, model.ErrNotFound) {
						return nil, xcode.RecordNotFound
					}
					l.Logger.Errorf("delete problem tag fail, err : %v, result : %+v", err, result)
				}
			}
		}
	}

	newData := &model.Problem{
		ProblemId: in.ProblemId,
		Title:     in.Title,
		Author:    in.Author,
		Des:       in.Des,
		Input:     in.Input,
		Output:    in.Output,
		Sinput:    in.Sinput,
		Soutput:   in.Soutput,
		Time:      in.Time,
		Memory:    in.Memory,
		//TestCount:   in.TestCount,
		Level:       in.Level,
		Oj:          in.Oj,
		Source:      in.Source,
		ProblemCode: in.ProblemCode,
		Hint:        in.Hint,
	}

	oneByTitle, err := l.svcCtx.ProblemModel.FindOneByTitle(l.ctx, newData.Title)
	if !errors.Is(err, model.ErrNotFound) && oneByTitle.ProblemId != in.ProblemId {
		return nil, code.ProblemTitleExist
	}

	err = l.svcCtx.ProblemModel.PartialUpdate(l.ctx, newData)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		l.Logger.Errorf("update problem fail, err : %v, newData : %+v", err, newData)
		return nil, xcode.ServerErr
	}
	return &pb.UpdateProblemResp{}, nil
}
