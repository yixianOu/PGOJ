package problems

import (
	"context"
	"encoding/json"
	"oj-micro/common/dataType"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/pb"
	"time"

	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProblemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProblemLogic {
	return &UpdateProblemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProblemLogic) UpdateProblem(req *types.UpdateProblemRequest) (resp *types.UpdateProblemResponse, err error) {
	//jwt是否过期
	expireTime, err := l.ctx.Value("access_expire").(json.Number).Int64()
	if err != nil || time.Now().Unix() > expireTime {
		return nil, xcode.UnauthorizedTokenTimeout
	}
	roleLevel, err := l.ctx.Value("role_level").(json.Number).Int64()
	userLevel := dataType.RoleLevel(roleLevel)
	if err != nil || userLevel != dataType.AdminUser {
		return nil, xcode.UnauthorizedUserNotSuperUser
	}

	_, err = l.svcCtx.ProblemServiceRpc.UpdateProblem(l.ctx, &pb.UpdateProblemReq{
		ProblemId: req.ProblemId,
		Author:    req.Author,
		Oj:        req.Oj,
		Title:     req.Title,
		Des:       req.Description,
		Input:     req.Input,
		Output:    req.Output,
		Sinput:    req.SampleInput,
		Soutput:   req.SampleOutput,
		Hint:      req.Hint,
		Source:    req.Source,
		Time:      req.LimitTime,
		Memory:    req.LimitMemory,
		Auth:      req.Auth,
		Level:     req.Level,
		//TestCount:    req.TestCount,
		ProblemCode:  req.ProblemCode,
		TagIds:       req.TagIds,
		TagOperation: req.TagOperation,
	})

	return
}
