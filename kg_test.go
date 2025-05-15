package kit

import (
	"fmt"
	//"github.com/micro-services-roadmap/atom-kit/gormx/source/g/dal"
	"testing"

	"github.com/micro-services-roadmap/atom-kit/kg"
)

// pk 值是零值就是灾难(如id=0)
func TestInit(t *testing.T) {
	Init("config_test.yaml")
	fmt.Printf("zap: %v", kg.L)

	//dal.SetDefault(kg.DB)
	//find, err := dal.Q.ArchivedUp.Preload(dal.Q.ArchivedUp.ArchivedUpsTag).Find()
	//if err != nil || find == nil {
	//	return
	//}
}
