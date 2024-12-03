package logic

import (
	"context"
	"database/sql"
	"errors"
	"oj-micro/common/xcode"
	"oj-micro/judgeStatus/model"

	"oj-micro/judgeStatus/cmd/rpc/internal/svc"
	"oj-micro/judgeStatus/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateJudgestatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateJudgestatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateJudgestatusLogic {
	return &UpdateJudgestatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateJudgestatusLogic) UpdateJudgestatus(in *pb.UpdateJudgestatusReq) (*pb.UpdateJudgestatusResp, error) {
	var inputData, userOutput, sampleOutPut sql.NullString
	if in.InputData != "" {
		inputData.String = in.InputData
		inputData.Valid = true
	}
	if in.UserOutput != "" {
		userOutput.String = in.UserOutput
		userOutput.Valid = true
	}
	if in.SampleOutput != "" {
		sampleOutPut.String = in.SampleOutput
		sampleOutPut.Valid = true
	}

	err := l.svcCtx.JudgeStatusModel.PartialUpdate(l.ctx, &model.Judgestatus{
		JudgeId:      in.JudgeId,
		Result:       in.Result,
		Time:         in.TimeCost,
		Memory:       in.MemoryCost,
		Testcase:     in.TestCases,
		Message:      in.Message,
		InputData:    inputData,
		UserOutPut:   userOutput,
		SampleOutPut: sampleOutPut,
	})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		l.Logger.Errorf("JudgeStatusModel PartialUpdate error: %v", err)
		return nil, xcode.ServerErr
	}
	return &pb.UpdateJudgestatusResp{}, nil
}
