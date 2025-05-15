package pb

import (
	"database/sql"
	"github.com/Mikaelemmmm/sql2pb/core"
	"github.com/alice52/jasypt-go"
	"github.com/micro-services-roadmap/atom-kit/kg"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

const (
	dbType     = "mysql"
	pkg        = "pb"
	goPkg      = "./pb"
	table      = "*"
	fieldStyle = "sql_pb"
)

// Deprecated: G0
func G0(service, dbType, mdsn string) error {
	if v, err := jasypt.New().Decrypt(mdsn); err == nil {
		mdsn = v
	}

	db, err := sql.Open(dbType, mdsn)
	if err != nil {
		return err
	}
	defer db.Close()

	return Gpb(db, service, pkg, goPkg, table, fieldStyle, nil, nil)
}

func G(service string) error {
	var dialector gorm.Dialector
	switch kg.C.System.DbType {
	case kg.DbMysql:
		dialector = mysql.Open(kg.C.Mysql.Dsn())
	case kg.DbPgsql:
		dialector = postgres.Open(kg.C.Pgsql.Dsn())
	default:
		panic("unknown db type")
	}

	gdb, err := gorm.Open(dialector)
	if err != nil {
		return err
	}
	db, err := gdb.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	return Gpb(db, service, pkg, goPkg, table, fieldStyle, nil, nil)
}

func Gpb(db *sql.DB, service, pkg, goPkg, table, fieldStyle string, ignoreTables, ignoreColumns []string) error {

	s, err := core.GenerateSchema(db, table, ignoreTables, ignoreColumns, service, goPkg, pkg, fieldStyle)
	if nil != err {
		return err
	}

	filePath := "./pb/" + service + ".proto"
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(filePath, []byte(s.String()), 0644); err != nil {
		return err
	}

	return nil
}
