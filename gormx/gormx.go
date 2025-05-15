package gormx

import (
	"fmt"
	"gorm.io/gen/field"
	"regexp"
	"strings"

	"github.com/micro-services-roadmap/atom-kit/gormx/initialize"
	"github.com/micro-services-roadmap/atom-kit/gormx/tenant"
	kg "github.com/micro-services-roadmap/atom-kit/kg"
	"gorm.io/gorm"
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

const (
	OrderAscending  = "ascending"
	OrderAsc        = "asc"
	OrderDescending = "descending"
	OrderDesc       = "desc"
)

func DynamicSort(tableName, sortBy, sortOrder string, defaultSort ...field.Expr) field.Expr {
	return DynamicSortAgg(tableName, sortBy, sortOrder, nil, defaultSort...)
}

// DynamicSortAgg
// order := DynamicSortAgg(q.TableName(), in.SortBy, in.SortOrder, func(f field.Field) field.Field { return f.Sum() }, q.ProductID.Desc())
func DynamicSortAgg(tableName, sortBy, sortOrder string,
	aggFunc func(field.Field) field.Field, // Function parameter for aggregation
	defaultSort ...field.Expr) field.Expr {
	if len(sortBy) > 0 && len(sortOrder) == 0 {
		sortOrder = OrderDescending
	}
	if len(sortBy) == 0 && len(sortOrder) > 0 {
		sortOrder = ""
	}

	var orderExpr field.Expr
	if len(sortBy) != 0 {
		sortField := field.NewField(tableName, CamelToSnake(sortBy))
		if aggFunc != nil {
			sortField = aggFunc(sortField)
		}
		if strings.EqualFold(sortOrder, OrderAscending) || strings.EqualFold(sortOrder, OrderAsc) {
			orderExpr = sortField.Asc()
		} else {
			orderExpr = sortField.Desc()
		}
	}

	if orderExpr != nil {
		return orderExpr
	}

	if len(defaultSort) != 0 {
		return defaultSort[0]
	}
	fmt.Println("sortBy and default sort is empty")

	return orderExpr // nil
}

// CamelToSnake 驼峰转换为下划线
func CamelToSnake(s string) string {
	re := regexp.MustCompile("([A-Z])")
	snake := re.ReplaceAllStringFunc(s, func(m string) string {
		return "_" + strings.ToLower(m)
	})

	return strings.TrimPrefix(snake, "_")
}

func Alias(fd field.Expr) string {
	return CamelToSnake(fd.ColumnName().String())
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
