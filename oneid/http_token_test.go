package oneid

import (
	"fmt"
	"github.com/micro-services-roadmap/atom-kit/kg"
	"testing"
)

func init() {
	kg.C.OneidConf.AuthenticationUrl = "http://localhost:8000/oneid/token"
	kg.C.OneidConf.AccessKeyId = "058d1d5f-a23f-47bc-8e9e-d31f412786ce"
	kg.C.OneidConf.AccessKeySecret = "972eff62-bd7c-4757-b332-b32b094a7aa5"
	kg.C.OneidConf.Subject = "wpp-admin"
}

func TestGetOneidToken(t *testing.T) {

	resp := GetOneidToken()
	fmt.Println(resp)

	resp2 := GetOneidToken("zack")
	fmt.Println(resp2)
}
