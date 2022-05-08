package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/8
 * @Desc: 非接口类型的方法集合
 */

func dumpMethodSet(i interface{}) {
	dynType := reflect.TypeOf(i)

	if dynType == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynType.NumMethod()
	if n == 0 {
		fmt.Printf("%s 's method set is empty!\n", dynType)
		return
	}

	fmt.Printf("%s 's method set :\n", dynType)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynType.Method(j).Name)
	}
	fmt.Printf("\n")
}

type T3 struct {
}

// 非接口类型的方法集合
func (T3) M11()  {}
func (T3) M21()  {}
func (*T3) M31() {}
func (*T3) M41() {}

func main() {
	//var n int
	//dumpMethodSet(n)
	//dumpMethodSet(&n)
	//
	//var t T3
	//dumpMethodSet(t)
	//dumpMethodSet(&t)

	// 思考题
	var t T4
	dumpMethodSet(t)
	dumpMethodSet(&t)

	var s S
	dumpMethodSet(s)
	dumpMethodSet(&s)
}

type T4 struct{}

func (T4) M1() {}
func (T4) M2() {}

type S T4
