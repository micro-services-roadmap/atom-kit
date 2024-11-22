package gz

import (
	"google.golang.org/grpc/status"
	"testing"
)

func TestUsage(t *testing.T) {

	var err error
	status.Convert(err).Code()
	status.Convert(err).Message()
	status.Convert(err).Details()

}
