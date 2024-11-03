package upload_file

import (
	"context"
	"encoding/json"
	"errors"
	"mime/multipart"
	"oj-micro/common/dataType"
	"oj-micro/common/xcode"
	"oj-micro/other/internal/code"
	"oj-micro/other/internal/svc/pkg"
	"oj-micro/problems/cmd/rpc/pb"
	"time"

	"oj-micro/other/internal/svc"
	"oj-micro/other/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTestCaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTestCaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTestCaseLogic {
	return &AddTestCaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTestCaseLogic) AddTestCase(req *types.AddTestCaseRequest) (resp *types.AddTestCaseResponse, err error) {
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

	//检查是否存在
	_, err = l.svcCtx.ProblemServiceRpc.GetTestcasesByProblemIdAndTestGroup(l.ctx, &pb.GetTestcasesByProblemIdAndTestGroupReq{
		ProblemId: req.ProblemId,
		TestGroup: req.TestGroup,
	})
	//存在则返回
	if err == nil {
		return nil, code.TestcaseExists
	}

	inputFile := l.ctx.Value("inputFile").(multipart.File)
	inputFileHeader := l.ctx.Value("inputFileHeader").(*multipart.FileHeader)
	outputFile := l.ctx.Value("outputFile").(multipart.File)
	outputFileHeader := l.ctx.Value("outputFileHeader").(*multipart.FileHeader)

	var inputFileKey, outputFileKey string
	for _, minioClient := range l.svcCtx.MinioClients {
		inputFileKey, err = pkg.PutFileToMinio(l.ctx, minioClient, pkg.InputFile, inputFile, inputFileHeader, uint64(req.ProblemId), l.svcCtx.Config)
		if err != nil {
			if errors.Is(err, code.FileExists) {
				return nil, code.FileExists
			}
			logx.Errorf("PutFileToOSS error：%v", err)
			return nil, xcode.ServerErr
		}

		outputFileKey, err = pkg.PutFileToMinio(l.ctx, minioClient, pkg.OutputFile, outputFile, outputFileHeader, uint64(req.ProblemId), l.svcCtx.Config)
		if err != nil {
			if errors.Is(err, code.FileExists) {
				return nil, code.FileExists
			}
			logx.Errorf("PutFileToOSS error：%v", err)
			return nil, xcode.ServerErr
		}
	}

	_, err = l.svcCtx.ProblemServiceRpc.AddTestcases(l.ctx, &pb.AddTestcasesReq{
		ProblemId:      req.ProblemId,
		TestGroup:      req.TestGroup,
		InputFileName:  inputFileKey,
		OutputFileName: outputFileKey,
	})
	if err != nil {
		return nil, err
	}

	return
}
