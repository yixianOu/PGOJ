package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"oj-micro/common/xcode"
)

// ServerErrorInterceptor 服务端错误拦截器, 将Error转换为Status
func ServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		return resp, xcode.StatusFromError(err).Err()
	}
}
