package tenant

import (
	"errors"
	"gorm.io/gorm"
)

type TenantHooks struct {
	TenantID string `gorm:"column:tenant_id;type:character varying(64);not null" json:"tenant_id"` // 标签名称
}

func (u *TenantHooks) GetTenantID(tx *gorm.DB) (string, error) {
	tenantID := GetTenantIDFromContext(tx.Statement.Context)
	if tenantID == "" {
		return "", errors.New("tenant_id is empty")
	}

	return tenantID, nil
}

func (u *TenantHooks) BeforeCreate(tx *gorm.DB) (err error) {
	tenantID, err := u.GetTenantID(tx)
	if err != nil {
		return err
	}
	u.TenantID = tenantID
	return
}

func (u *TenantHooks) BeforeSave(tx *gorm.DB) (err error) {
	tenantID, err := u.GetTenantID(tx)
	if err != nil {
		return err
	}
	u.TenantID = tenantID
	return
}

func (u *TenantHooks) BeforeUpdate(tx *gorm.DB) (err error) {
	return u.DoBeforeQuery(tx, err)
}

func (u *TenantHooks) BeforeDelete(tx *gorm.DB) (err error) {
	return u.DoBeforeQuery(tx, err)
}

func (u *TenantHooks) BeforeQuery(tx *gorm.DB) (err error) {
	return u.DoBeforeQuery(tx, err)
}

func (u *TenantHooks) DoBeforeQuery(tx *gorm.DB, err error) error {
	tenantID, err := u.GetTenantID(tx)
	if err != nil {
		return err
	}
	tx.Where("tenant_id = ?", tenantID)
	return nil
}
