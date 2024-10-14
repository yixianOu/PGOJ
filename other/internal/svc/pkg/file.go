package pkg

import (
	//"errors"

	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"oj-micro/other/internal/code"
	"oj-micro/other/internal/config"
	"path"
	"strconv"
	"strings"
	"time"
)

type FileType string

// 支持的文件请求类型
const (
	UserCover    FileType = "user_cover"
	ArticleImage FileType = "article_image"
	InputFile    FileType = "input_file"
	OutputFile   FileType = "output_file"
)

// GetFileName 1.如果是覆盖上传，则user_id+file_type+ext作为文件名。2.如果是追加上传，则MD5(file_name)+ext作为文件名
func GetFileName(fileName string, id uint64, fileType FileType) string {
	switch fileType {
	case UserCover:
		fileName = string(fileType) + "/" + strconv.FormatUint(id, 10) + GetFileExt(fileName)
	case ArticleImage:
		fileName = string(fileType) + "/" + strconv.FormatUint(id, 10) + "/" + time.Now().Format(time.UnixDate) + GetFileExt(fileName)
	case InputFile:
		fileName = strconv.FormatUint(id, 10) + "/" + string(fileType) + "/" + fileName
	case OutputFile:
		fileName = strconv.FormatUint(id, 10) + "/" + string(fileType) + "/" + fileName
	}

	return fileName
}

// GetFileExt 获取文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// CheckContainExt 检查文件后缀是否包含在约定的后缀配置项中
func CheckContainExt(t FileType, name string, config config.Config) bool {
	ext := GetFileExt(name)
	//统一格式
	ext = strings.ToUpper(ext)
	switch t {
	case UserCover:
		for _, allowExt := range config.UploadConfigs.ImageLimit.AllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	case ArticleImage:
		for _, allowExt := range config.UploadConfigs.ImageLimit.AllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	case InputFile:
		for _, allowExt := range config.UploadConfigs.SampleIOLimit.AllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	case OutputFile:
		for _, allowExt := range config.UploadConfigs.SampleIOLimit.AllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}

	return false
}

// OverMaxSize 检查文件大小是否超过大小限制
func OverMaxSize(t FileType, f *multipart.FileHeader, config config.Config) bool {
	switch t {
	case UserCover:
		if f.Size >= config.UploadConfigs.ImageLimit.MaxSize*1024*1024 {
			return true
		}
	case ArticleImage:
		if f.Size >= config.UploadConfigs.ImageLimit.MaxSize*1024*1024 {
			return true
		}
	case InputFile:
		if f.Size >= config.UploadConfigs.SampleIOLimit.MaxSize*1024*1024 {
			return true
		}
	case OutputFile:
		if f.Size >= config.UploadConfigs.SampleIOLimit.MaxSize*1024*1024 {
			return true
		}
	}

	return false
}

//
//func PutFileToOSS(c *oss.Client, fileType FileType, file multipart.File, header *multipart.FileHeader, userID uint64, config config.Config) (string, error) {
//	//文件检查(文件大小，后缀）
//	if !CheckContainExt(fileType, header.Filename, config) {
//		return "", errors.New("file suffix is not supported,fileType:" + string(fileType) + ",filename:" + header.Filename)
//	}
//	if OverMaxSize(fileType, header, config) {
//		return "", errors.New("exceeded maximum file limit")
//	}
//	//根据fileType生成文件名（包含用户信息）
//	fileName := GetFileName(header.Filename, userID, fileType)
//
//	//bucket, err := c.Bucket(config.MinioConfig.BucketName)
//	//if err != nil {
//	//	return "", err
//	//}
//	//
//	//err = bucket.PutObject(fileName, file)
//	if err != nil {
//		return "", err
//	}
//	return fileName, nil
//}

func PutFileToMinio(ctx context.Context, osClient *minio.Client, fileType FileType, file multipart.File, header *multipart.FileHeader, ID uint64, config config.Config) (string, error) {
	//文件检查(文件大小，后缀）
	if !CheckContainExt(fileType, header.Filename, config) {
		return "", errors.New("file suffix is not supported,fileType:" + string(fileType) + ",filename:" + header.Filename)
	}
	if OverMaxSize(fileType, header, config) {
		return "", errors.New("exceeded maximum file limit")
	}
	//根据fileType生成文件名（包含用户信息）
	fileName := GetFileName(header.Filename, ID, fileType)
	var err error

	switch fileType {
	case InputFile:
		_, err = osClient.StatObject(ctx, config.MinioConfig.BucketName, fileName, minio.StatObjectOptions{})
		if err == nil {
			return "", code.FileExists
		}
		_, err = osClient.PutObject(ctx, config.MinioConfig.BucketName, fileName, file, header.Size, minio.PutObjectOptions{
			ContentDisposition: "attachment;filename=" + header.Filename,
		})
	case OutputFile:
		_, err = osClient.StatObject(ctx, config.MinioConfig.BucketName, fileName, minio.StatObjectOptions{})
		if err == nil {
			return "", code.FileExists
		}
		_, err = osClient.PutObject(ctx, config.MinioConfig.BucketName, fileName, file, header.Size, minio.PutObjectOptions{
			ContentDisposition: "attachment;filename=" + header.Filename,
		})
	case UserCover:
		_, err = osClient.PutObject(ctx, config.MinioConfig.BucketName, fileName, file, header.Size, minio.PutObjectOptions{
			ContentDisposition: "inline",
		})
	case ArticleImage:
		_, err = osClient.StatObject(ctx, config.MinioConfig.BucketName, fileName, minio.StatObjectOptions{})
		if err == nil {
			return "", code.FileExists
		}
		_, err = osClient.PutObject(ctx, config.MinioConfig.BucketName, fileName, file, header.Size, minio.PutObjectOptions{
			ContentDisposition: "inline",
		})
	}
	if err != nil {
		return "", errors.New("PutObject error: " + err.Error())
	}

	return fileName, nil
}
