package fib

import (
	"testing"
)

// 斐波拉契数列
func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1
	// var (
	// 	a int = 1
	// 	b     = 1
	// )

	// 赋值可以进⾏行行⾃自动类型推断
	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

}

// 测试交换数据
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// 方式1:传统情况
	// tmp := a
	// a = b
	// b = tmp

	// 方式2:简化写法：在⼀一个赋值语句句中可以对多个变量量进⾏行行同时赋值
	a, b = b, a
	t.Log(a, b)
}
