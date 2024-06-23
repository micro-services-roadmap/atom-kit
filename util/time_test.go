package util

import (
	"fmt"
	"testing"
	"time"
)

func TestMustParse(t *testing.T) {
	ti := MustParse("2024-04-22T13:33:50.626Z")
	fmt.Println(ti)

	str := Format(ti)
	fmt.Println(str)

	str1 := Format((*time.Time)(nil), "2006-01-02 15:04:05")
	fmt.Println(str1)

	//fmt.Println(Format(nil)) // compile error
}

func TestParse(t *testing.T) {

}

func TestTryParse(t *testing.T) {

}
