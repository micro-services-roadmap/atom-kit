package tenant

import (
	"context"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestParseTenantIDFromMD(t *testing.T) {
	print(ParseTenantIDFromMD(metadata.New(nil)))
}

func TestWithTenantContext(t *testing.T) {
	ctx := WithTenantContext(context.Background(), metadata.New(nil))
	val := ctx.Value(TenantKey)
	print(val)
}
