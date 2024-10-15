package tags

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

type DeleteTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagLogic {
	return &DeleteTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTagLogic) DeleteTag(req *types.DeleteTagRequest) (resp *types.DeleteTagResponse, err error) {
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

	_, err = l.svcCtx.ProblemServiceRpc.DelTag(l.ctx, &pb.DelTagReq{
		Id: req.TagId,
	})
	if err != nil {
		return nil, err
	}
	return
}
