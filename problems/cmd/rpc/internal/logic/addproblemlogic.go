package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/internal/code"
	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"
	"oj-micro/problems/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProblemLogic {
	return &AddProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProblemLogic) AddProblem(in *pb.AddProblemReq) (*pb.AddProblemResp, error) {
	_, err := l.svcCtx.ProblemModel.FindOneByTitle(l.ctx, in.Title)
	if !errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorf("find problem fail, err : %v", err)
		return nil, code.ProblemTitleExist
	}
	if in.ProblemCode != "" {
		_, err := l.svcCtx.ProblemModel.FindOneByProblemCode(l.ctx, in.ProblemCode)
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.Errorf("find problem fail, err : %v", err)
			return nil, code.ProblemCodeExist
		}
	}

	result, err := l.svcCtx.ProblemModel.Insert(l.ctx, &model.Problem{
		Author:  in.Author, //jwt解析userID
		Oj:      in.Oj,
		Title:   in.Title,
		Des:     in.Des,
		Input:   in.Input,
		Output:  in.Output,
		Sinput:  in.Sinput,  //默认为空
		Soutput: in.Soutput, //默认为空
		Hint:    in.Hint,
		Source:  in.Source,
		Time:    in.Time,
		Memory:  in.Memory,
		Auth:    in.Auth,
		Level:   in.Level,
		//TestCount: in.TestCount, //默认为0
		ProblemCode: in.ProblemCode,
	})
	if err != nil {
		logx.Errorf("insert problem fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}

	problemId, err := result.LastInsertId()
	if err != nil {
		logx.Errorf("get problemId fail, err : %v", err)
		return nil, xcode.ServerErr
	}

	if in.ProblemCode == "" {
		err = l.svcCtx.ProblemModel.PartialUpdate(l.ctx, &model.Problem{
			ProblemId:   problemId,
			ProblemCode: "p" + fmt.Sprintf("%04d", problemId),
		})
		if err != nil {
			logx.Errorf("update problem fail, err : %v", err)
			return nil, xcode.ServerErr
		}
	}

	result, err = l.svcCtx.ProblemdataModel.Insert(l.ctx, &model.Problemdata{
		ProblemId:  problemId,
		Score:      in.Score,
		Auth:       in.Auth,
		Submission: 0,
		Ac:         0,
		Mle:        0,
		Tle:        0,
		Wa:         0,
		Rte:        0,
		Ce:         0,
		Ole:        0,
		Ue:         0,
		Sf:         0,
		Fe:         0,
	})
	if err != nil {
		logx.Errorf("insert problemdata fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}

	for _, tag := range in.TagIds {
		result, err = l.svcCtx.ProblemTagModel.Insert(l.ctx, &model.ProblemTag{
			ProblemId: problemId,
			TagId:     tag,
		})
		if err != nil {
			logx.Errorf("insert problemtag fail, err : %v, result : %+v", err, result)
			return nil, xcode.ServerErr
		}
	}
	return &pb.AddProblemResp{}, nil
}
