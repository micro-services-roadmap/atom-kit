package g

import (
	"fmt"
	"github.com/micro-services-roadmap/kit-common/gormx/tenant/t"

	"github.com/alice52/jasypt-go"
	"github.com/micro-services-roadmap/kit-common/kg"
	ggy "github.com/we7coreteam/gorm-gen-yaml"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

var (
	autoUpdateTimeField = gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
		return map[string][]string{"column": {"update_time"}, "type": {"datetime(3)"}, "autoUpdateTime": {}}
	})
	autoCreateTimeField = gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
		return map[string][]string{"column": {"create_time"}, "type": {"datetime(3)"}, "autoCreateTime": {}}
	})
	softDeleteField = gen.FieldType("delete_time", "gorm.DeletedAt")
)
var FieldOpts = []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField, softDeleteField}

// https://zhuanlan.zhihu.com/p/653483236
// https://github.com/Alice52/go-tutorial/issues/5#issuecomment-1286325129

// Deprecated: use G2
func G(dbTpe, dsn, outputDir, relationYaml string, opts ...gen.ModelOpt) (*gen.Generator, *gorm.DB) {
	var DSN string
	if v, err := jasypt.New().Decrypt(dsn); err != nil {
		DSN = dsn
	} else {
		DSN = v
	}
	var dialector gorm.Dialector
	switch dbTpe {
	case kg.DbMysql:
		dialector = mysql.Open(DSN)
	case kg.DbPgsql:
		dialector = postgres.Open(DSN)
	default:
		panic("unknown db type")
	}

	return genCore(dialector, outputDir, relationYaml, opts...)
}

func G2(outputDir, relationYaml string, useTenant bool, opts ...gen.ModelOpt) (*gen.Generator, *gorm.DB) {
	var dialector gorm.Dialector
	switch kg.C.System.DbType {
	case kg.DbMysql:
		dialector = mysql.Open(kg.C.Mysql.Dsn())
	case kg.DbPgsql:
		dialector = postgres.Open(kg.C.Pgsql.Dsn())
	default:
		panic("unknown db type")
	}

	if useTenant {
		opts = append(opts, gen.WithMethod(t.TenantHooks{}))
	}

	return genCore(dialector, outputDir, relationYaml, opts...)
}

func genCore(dialector gorm.Dialector, outputDir string, relationYaml string, opts ...gen.ModelOpt) (*gen.Generator, *gorm.DB) {
	// 连接数据库
	db, err := gorm.Open(dialector)
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           outputDir, //"./source/gen/dal",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db)

	var columOfInt = func(columnType gorm.ColumnType) (dataType string) {
		if n, ok := columnType.Nullable(); ok && n {
			return "*int64"
		}
		return "int64"
	}
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   columOfInt,
		"smallint":  columOfInt,
		"mediumint": columOfInt,
		"bigint":    columOfInt,
		"int":       columOfInt,
		"int2":      columOfInt, // pgsql type
		"int4":      columOfInt, // pgsql type
		"int8":      columOfInt, // pgsql type
	}

	g.WithDataTypeMap(dataMap)
	g.WithOpts(append(FieldOpts, opts...)...)

	ggy.NewYamlGenerator(relationYaml).UseGormGenerator(g).Generate(FieldOpts...) // fieldOpts is not used
	//g.ApplyBasic(g.GenerateAllTable()...) // will lose relation, so donot use it after NewYamlGenerator
	//g.ApplyInterface(func(upsTagInterface) {}, g.GenerateModel("archived_ups_tag", FieldOpts...))

	return g, db
}
