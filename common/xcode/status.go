package xcode

import (
	"context"
	"github.com/golang/protobuf/ptypes"

	//"errors"
	"fmt"
	"oj-micro/common/xcode/types"
	"strconv"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var _ XCode = (*Status)(nil)

// Status 组合types.Status结构体，实现XCode接口
type Status struct {
	sts *types.Status
}

func Error(code Code) *Status {
	return &Status{sts: &types.Status{Code: int32(code.Code()), Message: code.Message()}}
}

func Errorf(code Code, format string, args ...interface{}) *Status {
	code.msg = fmt.Sprintf(format, args...)
	return Error(code)
}

func (s *Status) Error() string {
	return s.Message()
}

func (s *Status) Code() int {
	return int(s.sts.Code)
}

func (s *Status) Message() string {
	if s.sts.Message == "" {
		return strconv.Itoa(int(s.sts.Code))
	}

	return s.sts.Message
}

func (s *Status) Details() []interface{} {
	if s == nil || s.sts == nil {
		return nil
	}
	details := make([]interface{}, 0, len(s.sts.Details))
	for _, d := range s.sts.Details {
		detail := &ptypes.DynamicAny{}

		if err := d.UnmarshalTo(detail); err != nil {
			details = append(details, err)
			continue
		}
		details = append(details, detail.Message)
	}

	return details
}

func (s *Status) WithDetails(msgs ...proto.Message) (*Status, error) {
	for _, msg := range msgs {
		anyMsg, err := anypb.New(msg)
		if err != nil {
			return s, err
		}
		s.sts.Details = append(s.sts.Details, anyMsg)
	}

	return s, nil
}

func (s *Status) Proto() *types.Status {
	return s.sts
}

// FromCodeToStatus 将错误码Code结构体转换为grpc状态码Status结构体
func FromCodeToStatus(code Code) *Status {
	return &Status{sts: &types.Status{Code: int32(code.Code()), Message: code.Message()}}
}

// FromProtoToCode 将proto.Message转换为Code或Status并返回,满足XCode接口
func FromProtoToCode(pbMsg proto.Message) XCode {
	msg, ok := pbMsg.(*types.Status)
	// 如果msg是Status类型
	if ok {
		// 如果message为空或者message等于code，则返回Code
		if len(msg.Message) == 0 || msg.Message == strconv.FormatInt(int64(msg.Code), 10) {
			return Code{code: int(msg.Code)}
		}
		// 否则返回携带完整信息的Status（code+message）
		return &Status{sts: msg}
	}

	return Errorf(ServerErr, "invalid proto message get %v", pbMsg)
}

// statusToXCode 将grpc的status.Status转换为项目错误码Code结构体
func statusToXCode(grpcStatus *status.Status) Code {
	grpcCode := grpcStatus.Code()
	switch grpcCode {
	case codes.OK:
		return OK
	case codes.InvalidArgument:
		return RequestErr
	case codes.NotFound:
		return NotFound
	case codes.PermissionDenied:
		return AccessDenied
	case codes.Unauthenticated:
		return Unauthorized
	case codes.ResourceExhausted:
		return LimitExceed
	case codes.Unimplemented:
		return MethodNotAllowed
	case codes.DeadlineExceeded:
		return Deadline
	case codes.Unavailable:
		return ServiceUnavailable
	case codes.Unknown:
		return String(grpcStatus.Message())
	}

	return ServerErr
}

func CodeFromError(err error) XCode {
	//在最后返回结果的时候才解析根错误
	err = errors.Cause(err)
	if code, ok := err.(XCode); ok {
		return code
	}

	switch err {
	case context.Canceled:
		return Canceled
	case context.DeadlineExceeded:
		return Deadline
	}

	return ServerErr
}

// StatusFromError 将error转换为grpc的status.Status
func StatusFromError(err error) *status.Status {
	err = errors.Cause(err)
	if code, ok := err.(XCode); ok {
		grpcStatus, e := gRPCStatusFromXCode(code)
		if e == nil {
			return grpcStatus
		}
	}

	var grpcStatus *status.Status
	switch err {
	case context.Canceled:
		grpcStatus, _ = gRPCStatusFromXCode(Canceled)
	case context.DeadlineExceeded:
		grpcStatus, _ = gRPCStatusFromXCode(Deadline)
	default:
		grpcStatus, _ = status.FromError(err)
	}

	return grpcStatus
}

// gRPCStatusFromXCode 将XCode接口转换为grpc的status.Status
func gRPCStatusFromXCode(code XCode) (*status.Status, error) {
	var sts *Status
	switch v := code.(type) {
	case *Status:
		sts = v
	case Code:
		sts = FromCodeToStatus(v)
	default:
		sts = Error(Code{code.Code(), code.Message()})
		for _, detail := range code.Details() {
			if msg, ok := detail.(proto.Message); ok {
				_, _ = sts.WithDetails(msg)
			}
		}
	}

	stas := status.New(codes.Unknown, strconv.Itoa(sts.Code()))
	return stas.WithDetails(sts.Proto())
}

// GrpcStatusToXCode 将grpc的status.Status转换为微服务错误码XCode接口
func GrpcStatusToXCode(gstatus *status.Status) XCode {
	details := gstatus.Details()
	//遍历details，找到最后一个proto.Message类型的detail
	for i := len(details) - 1; i >= 0; i-- {
		detail := details[i]
		//如果detail是proto.Message类型，则转换为XCode
		if pb, ok := detail.(proto.Message); ok {
			return FromProtoToCode(pb)
		}
	}
	//如果没有找到proto.Message类型的detail，则直接返回gstatus的code对应的业务错误码
	return statusToXCode(gstatus)
}
