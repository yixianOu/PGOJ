package upload_file

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"oj-micro/common/xcode"
	"oj-micro/other/internal/svc/pkg"
	"oj-micro/users/cmd/rpc/pb"
	"strconv"
	"time"

	"oj-micro/other/internal/svc"
	"oj-micro/other/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserCoverLogic {
	return &UpdateUserCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserCoverLogic) UpdateUserCover(req *types.UpdateUserCoverRequest) (resp *types.UpdateUserResponse, err error) {
	//jwt是否过期
	expireTime, err := l.ctx.Value("access_expire").(json.Number).Int64()
	if err != nil || time.Now().Unix() > expireTime {
		return nil, xcode.UnauthorizedTokenTimeout
	}
	userID, err := l.ctx.Value("user_id").(json.Number).Int64()
	if err != nil || userID != req.ID {
		return nil, xcode.UnauthorizedUserNotLogin
	}

	file := l.ctx.Value("CoverFile").(multipart.File)
	fileHeader := l.ctx.Value("CoverFileHeader").(*multipart.FileHeader)

	//删除原来的用户头像
	done := l.svcCtx.MinioClients[0].ListObjects(l.ctx, l.svcCtx.Config.MinioConfig.BucketName, minio.ListObjectsOptions{
		Prefix:  string(pkg.UserCover) + "/" + strconv.FormatUint(uint64(userID), 10),
		MaxKeys: 1,
	})
	removeObjectErrors := l.svcCtx.MinioClients[0].RemoveObjects(l.ctx, l.svcCtx.Config.MinioConfig.BucketName, done, minio.RemoveObjectsOptions{})
	info := <-removeObjectErrors
	if info.Err != nil {
		logx.Errorf("remove object fail, err : %v", info.Err)
		return nil, xcode.ServerErr
	}

	fileKey, err := pkg.PutFileToMinio(l.ctx, l.svcCtx.MinioClients[0], pkg.UserCover, file, fileHeader, uint64(userID), l.svcCtx.Config)
	if err != nil {
		logx.Errorf("PutFileToOSS error：%v", err)
		return nil, xcode.ServerErr
	}

	//预签名
	//CoverUrl, err := l.svcCtx.MinioClients.PresignedGetObject(l.ctx, l.svcCtx.Config.MinioConfig.BucketName, fileKey, time.Hour*24*365, url.Values{
	//	"response-content-disposition": {"inline"},
	//})
	//if err != nil {
	//	logx.Errorf("PresignedGetObject error：%v", err)
	//	return nil, xcode.ServerErr
	//}
	CoverUrl := fmt.Sprintf("%s/%s/%s", l.svcCtx.Config.MinioConfig.DomainName, l.svcCtx.Config.MinioConfig.BucketName, fileKey)
	//log.Printf("CoverUrl: %s", CoverUrl)
	_, err = l.svcCtx.UserServiceRpc.PartialUpdateUser(l.ctx, &pb.PartialUpdateUserReq{
		Id:            userID,
		CoverImageUrl: CoverUrl,
	})

	return nil, err
}
