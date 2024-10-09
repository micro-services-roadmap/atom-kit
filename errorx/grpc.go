package errorx

import (
	"github.com/micro-services-roadmap/oneid-core/model"
	"google.golang.org/grpc/status"
)

func GrpcError(err error) error {
	if err == nil {
		return nil
	}

	ge := status.Convert(err)
	return model.NewError(int(ge.Code()), ge.Message(), err)
}
