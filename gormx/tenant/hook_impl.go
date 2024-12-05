package tenant

import (
	"errors"
	"gorm.io/gorm"
)

func GetTenantID(tx *gorm.DB) (string, error) {
	tenantID := GetTenantIDFromContext(tx.Statement.Context)
	if tenantID == "" {
		return "", errors.New("tenant_id is empty")
	}

	return tenantID, nil
}

func DoBeforeQuery(tx *gorm.DB, err error) error {
	tenantID, err := GetTenantID(tx)
	if err != nil {
		return err
	}
	tx.Where("tenant_id = ?", tenantID)
	return nil
}
