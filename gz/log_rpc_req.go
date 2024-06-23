package gz

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	logx.WithContext(ctx).Debugf("[RPC-SRV] receive request: {%v}; method: [%+v]", req, info.FullMethod)
	resp, err = handler(ctx, req)
	logx.WithContext(ctx).Debugf("[RPC-SRV] respond request: {%v} error: [%v]", resp, err)

	return resp, err
}
