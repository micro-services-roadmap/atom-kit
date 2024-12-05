package t

import (
	"github.com/micro-services-roadmap/kit-common/gormx/tenant"
	"gorm.io/gorm"
)

type TenantHooks struct {
	TenantID string `gorm:"column:tenant_id;type:character varying(64);not null" json:"tenant_id"` // 标签名称
}

func (u *TenantHooks) BeforeCreate(tx *gorm.DB) (err error) {
	tenantID, err := tenant.GetTenantID(tx)
	if err != nil {
		return err
	}
	u.TenantID = tenantID
	return
}

func (u *TenantHooks) BeforeSave(tx *gorm.DB) (err error) {
	tenantID, err := tenant.GetTenantID(tx)
	if err != nil {
		return err
	}
	u.TenantID = tenantID
	return
}

func (u *TenantHooks) BeforeUpdate(tx *gorm.DB) (err error) {
	return tenant.DoBeforeQuery(tx, err)
}

func (u *TenantHooks) BeforeDelete(tx *gorm.DB) (err error) {
	return tenant.DoBeforeQuery(tx, err)
}

func (u *TenantHooks) BeforeQuery(tx *gorm.DB) (err error) {
	return tenant.DoBeforeQuery(tx, err)
}
