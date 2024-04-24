package gormx

import (
	"fmt"
	"testing"

	kg "github.com/micro-services-roadmap/kit-common/kg"
	"github.com/micro-services-roadmap/kit-common/viperx"
)

func init() {
	viperx.InitViperV0("config.yaml")
}

func TestInitDB(t *testing.T) {
	kg.C.System.DbType = kg.DbMysql
	// 测试InitDB函数
	InitDB()
	fmt.Printf("mysql: %v", kg.DB)
}
