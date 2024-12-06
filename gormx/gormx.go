package gormx

import (
	"github.com/micro-services-roadmap/kit-common/gormx/initialize"
	"github.com/micro-services-roadmap/kit-common/gormx/tenant"
	kg "github.com/micro-services-roadmap/kit-common/kg"
	"gorm.io/gorm"
	"strings"
)

var Unscoped = func(db *gorm.DB) *gorm.DB { return db.Unscoped() }

func ILike(column *string) string {
	if column == nil {
		return ""
	}

	return ILikeHelper(*column)
}

func ILikeHelper(column string) string {
	if len(column) == 0 {
		return ""
	}

	return "%" + strings.ToLower(column) + "%"
}

func Page[T int | int8 | int16 | int32 | int64](current, pageSize T) (offset, limit int) {

	offset = int((current - 1) * pageSize)
	limit = int(pageSize)

	return
}

func InitDB() *gorm.DB {
	var useTenant bool
	dbType := kg.C.System.DbType
	switch dbType {
	case kg.DbMysql:
		kg.DB = initialize.GormMysql(kg.C.Mysql.Migration)
		useTenant = kg.C.Mysql.UseTenant
	case kg.DbPgsql:
		kg.DB = initialize.GormPgSQL(kg.C.Pgsql.Migration)
		useTenant = kg.C.Pgsql.UseTenant
	default:
		panic("unknown db type")
	}

	if useTenant {
		tenant.RegisterBeforeQuery(kg.DB)
	}

	return kg.DB
}
