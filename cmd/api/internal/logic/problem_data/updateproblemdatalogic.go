package problem_data

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

type UpdateProblemDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProblemDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProblemDataLogic {
	return &UpdateProblemDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProblemDataLogic) UpdateProblemData(req *types.UpdateProblemDataRequest) (resp *types.UpdateProblemDataResponse, err error) {
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

	_, err = l.svcCtx.ProblemServiceRpc.UpdateProblemdata(l.ctx, &pb.UpdateProblemdataReq{
		ProblemdataId: req.ProblemDataId,
		//Submission:    req.Submission,
		//Ac:            req.Accepted,
		//Mle:           req.MemoryLimitExceeded,
		//Tle:           req.TimeLimitExceeded,
		//Wa:            req.WrongAnswer,
		//Rte:           req.RuntimeError,
		//Ce:            req.CompileError,
		//Pe:            req.PresentationError,
		//Se:            req.SystemError,
		Score: req.Score,
		Auth:  req.Auth,
	})
	return
}
