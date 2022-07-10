package ch3

import (
	"fmt"
	"testing"
)

const IPv4Len = 4

func TestConst(t *testing.T) {
	s := "hello,const"
	parseIPv4(s)
}

func parseIPv4(s string) {
	var p [IPv4Len]byte
	/// do something
	fmt.Println(p)
}
