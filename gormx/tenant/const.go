package tenant

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const TenantKey = "X-Tenant-ID"
const TenantColumn = "tenant_id"

func ParseTenantIDFromMD(md metadata.MD) *string {
	vals := ParseFromMD(TenantKey, md)
	if len(vals) == 0 {
		return nil
	}

	return &vals[0]
}

func ParseFromMD(key string, md metadata.MD) []string {
	val, ok := md[key]
	if !ok {
		return nil
	}
	return val
}

func WithTenantContext(ctx context.Context, md metadata.MD) context.Context {
	return context.WithValue(ctx, TenantKey, ParseTenantIDFromMD(md))
}
