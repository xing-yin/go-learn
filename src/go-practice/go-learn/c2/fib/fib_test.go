package fib

import (
	"fmt"
	"testing"
)

// 斐波拉契数列
func TestFib(t *testing.T) {
	var a int = 1
	var b int = 2

	var (
		c int = 1
		d     = 2
	)
	fmt.Println(a, b, c, d)

	e := 1
	f := 1
	t.Log(e)
	// 1,1,2,3,5
	for i := 0; i < 5; i++ {
		t.Log(" ", f)
		tmp := e
		e = f
		f = tmp + e
	}
	fmt.Println(f)
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// 历史做法1
	tmp := a
	a = b
	b = tmp
	t.Log(a, b)

	// 方式2
	a, b = b, a
	t.Log(a, b)
}
