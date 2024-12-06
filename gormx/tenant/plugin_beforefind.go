package tenant

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
)

func RegisterBeforeQuery(db *gorm.DB) {
	err := db.Callback().Query().Before("gorm:query").Register("ltc:before_query", BeforeQuery)
	if err != nil {
		fmt.Printf("Register BeforeQuery Failed: %v", err)
	} else {
		fmt.Println("Register BeforeQuery Success")
	}
}

var (
	iface = reflect.TypeOf((*BeforeQueryInterface)(nil)).Elem()
	hook  = &TenantHooks{}
)

type BeforeQueryInterface interface {
	BeforeQuery(db *gorm.DB) error
}

func BeforeQuery(db *gorm.DB) {
	if db.Error == nil && !db.Statement.SkipHooks {
		value := db.Statement.ReflectValue.Interface()
		rt := reflect.TypeOf(value)
		// Helper function to execute the hook
		executeHook := func() {
			db.AddError(DoBeforeQuery(db))
		}

		switch rt.Kind() {
		case reflect.Slice, reflect.Array:
			elemType := rt.Elem()
			if elemType.Kind() == reflect.Ptr && elemType.Implements(iface) {
				executeHook()
			}
			if elemType.Kind() == reflect.Struct && reflect.PointerTo(elemType).Implements(iface) {
				executeHook()
			}
		case reflect.Ptr:
			if rt.Implements(iface) {
				executeHook()
			}
		case reflect.Struct:
			if reflect.PointerTo(rt).Implements(iface) {
				executeHook()
			}
		}
	}
}

func DoBeforeQuery(tx *gorm.DB) error {
	tenantID, err := hook.GetTenantID(tx)
	if err != nil {
		return err
	}

	tx.Where(TenantColumn+" = ?", *tenantID)
	return nil
}
