package ipx

import (
	"fmt"
	"github.com/gookit/goutil/fmtutil"
	"testing"
)

func TestQueryFreeByIPInfo(t *testing.T) {

	got, err := QueryByIPInfo("2601:0243:21:1448:da65:6d64:23d7:f37c", "5b18fa715132d6")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmtutil.PrettyJSON(got))
}
