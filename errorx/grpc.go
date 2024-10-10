package errorx

import (
	"github.com/micro-services-roadmap/oneid-core/model"
	"google.golang.org/grpc/status"
)

func GrpcError(err error) error {
	if err == nil {
		return nil
	}

	if s, ok := status.FromError(err); ok {
		return model.NewError(int(s.Code()), s.Message(), err)
	} else {
		return err
	}
}

func IsGrpcError(err error) bool {
	if err == nil {
		return false
	}

	_, ok := status.FromError(err)
	return ok
}
