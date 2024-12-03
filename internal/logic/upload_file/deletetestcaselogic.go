package upload_file

import (
	"context"
	"encoding/json"
	"github.com/minio/minio-go/v7"
	"oj-micro/common/dataType"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/pb"
	"time"

	"oj-micro/other/internal/svc"
	"oj-micro/other/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTestCaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTestCaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTestCaseLogic {
	return &DeleteTestCaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTestCaseLogic) DeleteTestCase(req *types.DeleteTestCaseRequest) (resp *types.DeleteTestCaseResponse, err error) {
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

	testcasesById, err := l.svcCtx.ProblemServiceRpc.GetTestcasesById(l.ctx, &pb.GetTestcasesByIdReq{
		Id: req.TestId,
	})
	if err != nil {
		return nil, err
	}

	for _, minioClient := range l.svcCtx.MinioClients {
		err = minioClient.RemoveObject(l.ctx, l.svcCtx.Config.MinioConfig.BucketName, testcasesById.Testcases.InputFileName, minio.RemoveObjectOptions{
			ForceDelete: true,
		})
		if err != nil {
			l.Logger.Errorf("remove object fail, err : %v", err)
			return nil, xcode.ServerErr
		}
		err = minioClient.RemoveObject(l.ctx, l.svcCtx.Config.MinioConfig.BucketName, testcasesById.Testcases.OutputFileName, minio.RemoveObjectOptions{
			ForceDelete: true,
		})
		if err != nil {
			l.Logger.Errorf("remove object fail, err : %v", err)
			return nil, xcode.ServerErr
		}
	}

	_, err = l.svcCtx.ProblemServiceRpc.DelTestcases(l.ctx, &pb.DelTestcasesReq{
		Id: req.TestId,
	})
	if err != nil {
		return nil, err
	}
	return
}
