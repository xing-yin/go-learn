package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeConvert(t *testing.T) {
	var sum int = 17
	var count int = 5
	var mean float32

	// go 不支持隐式转换类型: Cannot use 'sum / count' (type int) as the type float32
	// mean = sum / count

	mean = float32(sum) / float32(count)
	fmt.Printf("mean value is:%f", mean)
}

type Handler func()

func a() Handler {
	return func() {
	}
}

func Test1(t *testing.T) {
	var i interface{} = testN
	_, ok := i.(func())
	fmt.Println(ok)

	_, ok = i.(Handler)
	fmt.Println(ok)
	fmt.Println(reflect.TypeOf(testN) == reflect.TypeOf((*Handler)(nil)).Elem())
}
