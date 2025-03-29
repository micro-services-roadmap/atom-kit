package gz

import (
	"context"
	"github.com/kongmsr/oneid-core/modelo"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	logx.WithContext(ctx).Debugf("[RPC-SRV] receive request: {%v}; method: [%+v]", req, info.FullMethod)
	resp, err = handler(ctx, req)
	logx.WithContext(ctx).Debugf("[RPC-SRV] respond request: {%v} error: [%v]", resp, err)

	if mcer, ok := err.(*modelo.CodeError); !ok {
		return resp, err
	} else if s, ok := status.FromError(err); ok {
		return resp, s.Err()
	} else {
		var msg string
		if len(mcer.Msg) == 0 && mcer.Err != nil {
			msg = mcer.Err.Error()
		} else {
			msg = mcer.Msg
		}
		return resp, status.Error(gcodes.Code(mcer.Code), msg)
	}
}
