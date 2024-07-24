package g

import (
	"github.com/micro-services-roadmap/kit-common/kg"
	"gorm.io/gen"
	"testing"
)

var pgsqlDsn = "host=xxx port=5431 user=zack password=xxx dbname=kit-common sslmode=disable TimeZone=Asia/Shanghai"

func TestG(t *testing.T) {
	cnt_info := gen.FieldType("cnt_info", "datatypes.JSONType[CntInfo]")
	g, _ := G(kg.DbPgsql, pgsqlDsn, "./dal", "./relation.yaml", cnt_info)
	//info := g.Data[db.NamingStrategy.SchemaName("archived_ups_tag")]
	//g.ApplyInterface(func(upsTagInterface) {}, info.QueryStructMeta /*g.GenerateModel("archived_ups_tag")*/) /* relation will lose*/
	g.ApplyInterface(func(upsTagInterface) {}, g.GenerateModel("archived_fav", FieldOpts...)) /* relation will lose*/
	g.Execute()

}
