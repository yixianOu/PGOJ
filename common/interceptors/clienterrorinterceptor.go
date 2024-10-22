package interceptors

import (
	"context"
	"oj-micro/common/xcode"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ClientErrorInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// 调用实际的 gRPC 方法，传递请求和响应参数，并且捕获返回的错误
		err := invoker(ctx, method, req, reply, cc, opts...)
		// 如果调用过程发生了错误
		if err != nil {
			// 从错误中提取 gRPC status
			grpcStatus, _ := status.FromError(err)
			// 将 gRPC status 转换为自定义的 xcode 错误码 (xc)
			xc := xcode.GrpcStatusToXCode(grpcStatus)
			// 使用带有原始错误消息的自定义错误码构建一个新的错误
			err = errors.WithMessage(xc, grpcStatus.Message())
		}
		// 返回处理后的错误（或无错误）
		return err
	}
}
