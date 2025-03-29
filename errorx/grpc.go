package errorx

import (
	"github.com/kongmsr/oneid-core/modelo"
	"google.golang.org/grpc/status"
)

func GrpcError(err error) error {
	if err == nil {
		return nil
	}

	if s, ok := status.FromError(err); ok {
		return modelo.NewError(int(s.Code()), s.Message(), err)
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
