package gormx

import (
	"github.com/micro-services-roadmap/kit-common/gormx/initialize"
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

func Page[T int | int8 | int16 | int32 | int64](current, pageSize T) (offset, limit T) {

	offset = (current - 1) * pageSize
	limit = pageSize

	return
}

func InitDB() *gorm.DB {
	dbType := kg.C.System.DbType
	switch dbType {
	case kg.DbMysql:
		kg.DB = initialize.GormMysql(kg.C.Mysql.Migration)
	case kg.DbPgsql:
		kg.DB = initialize.GormPgSQL(kg.C.Pgsql.Migration)
	default:
		panic("unknown db type")
	}

	return kg.DB
}
