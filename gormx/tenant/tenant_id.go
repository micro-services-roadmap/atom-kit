package tenant

import "context"

// GetTenantIDFromContext 从上下文中获取 tenant_id
func GetTenantIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	tenantID, ok := ctx.Value("X-Tenant-ID").(string)
	if !ok {
		return ""
	}
	return tenantID
}
