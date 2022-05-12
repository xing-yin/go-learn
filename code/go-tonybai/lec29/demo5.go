package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/11
 * @Desc:
 */

type NonEmptyInterface interface {
	M1()
	M2()
}

type T8 struct {
	n int
	s string
}

func (T8) M1() {}

func (T8) M2() {}

func main() {
	var t = T8{
		n: 18,
		s: "hello,iface",
	}

	var i NonEmptyInterface = t
	fmt.Printf("%T\n", i)
}
