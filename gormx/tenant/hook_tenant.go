package tenant

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TenantHooks struct {
	TenantID string `gorm:"column:tenant_id;type:character varying(64);not null" json:"tenant_id"` // 标签名称
}

func (u *TenantHooks) BeforeCreate(tx *gorm.DB) (err error) {
	tenantID, err := u.GetTenantID(tx)
	if err != nil {
		return err
	}
	u.TenantID = *tenantID
	return
}

func (u *TenantHooks) BeforeSave(tx *gorm.DB) (err error) {
	tenantID, err := u.GetTenantID(tx)
	if err != nil {
		return err
	}
	u.TenantID = *tenantID
	return
}

func (u *TenantHooks) BeforeUpdate(tx *gorm.DB) (err error) {
	return u.DoBeforeQuery(tx, err)
}

func (u *TenantHooks) BeforeDelete(tx *gorm.DB) (err error) {
	return u.DoBeforeQuery(tx, err)
}

func (u *TenantHooks) DoBeforeQuery(tx *gorm.DB, err error) error {
	tenantID, err := u.GetTenantID(tx)
	if err != nil {
		return err
	}

	// 动态添加查询条件到 Statement
	tx.Statement.AddClause(clause.Where{
		Exprs: []clause.Expression{
			clause.Eq{Column: u.TenantColumn, Value: tenantID},
		},
	})

	return nil
}

func (u *TenantHooks) GetTenantID(tx *gorm.DB) (*string, error) {
	if tx.Statement.Context == nil {
		return nil, errors.New("no context found")
	}
	tenantID, ok := tx.Statement.Context.Value(u.TenantKey).(*string)
	if !ok || len(*tenantID) == 0 {
		return nil, errors.New("tenant_id is empty")
	}
	return tenantID, nil
}

func (u *TenantHooks) TenantKey(tx *gorm.DB) string {

	return "X-Tenant-ID"
}

func (u *TenantHooks) TenantColumn(tx *gorm.DB) string {
	return "tenant_id"
}
