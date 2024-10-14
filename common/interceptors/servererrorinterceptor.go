package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"oj-micro/common/xcode"
)

func ServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		return resp, xcode.StatusFromError(err).Err()
	}
}
