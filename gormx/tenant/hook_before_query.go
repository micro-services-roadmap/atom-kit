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

type BeforeQueryInterface interface {
	BeforeQuery(db *gorm.DB) error
}

func BeforeQuery(db *gorm.DB) {
	if db.Error == nil {
		callMethod(db, func(value interface{}, tx *gorm.DB) (called bool) {
			if i, ok := value.(BeforeQueryInterface); ok {
				called = true
				db.AddError(i.BeforeQuery(tx))
			}
			return called
		})
	}
}

func callMethod(db *gorm.DB, fc func(value interface{}, tx *gorm.DB) bool) {
	tx := db.Session(&gorm.Session{})
	if called := fc(db.Statement.Dest, tx); !called {
		switch db.Statement.ReflectValue.Kind() {
		case reflect.Slice, reflect.Array:
			db.Statement.CurDestIndex = 0
			for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
				fc(reflect.Indirect(db.Statement.ReflectValue.Index(i)).Addr().Interface(), tx)
				db.Statement.CurDestIndex++
			}
		case reflect.Struct:
			fc(db.Statement.ReflectValue.Addr().Interface(), tx)
		}
	}
}
