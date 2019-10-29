package customer_type

import (
	"fmt"
	"testing"
	"time"
)

// 原函数
//func timeSpent(inner func(op int) int) func(op int) int {
//	return func(n int) int {
//		start := time.Now()
//		// 这是真正运行的函数，有点想装饰函数
//		ret := inner(n)
//		fmt.Println("time spent:", time.Since(start).Seconds())
//		return ret
//	}
//}

// 自定义类型,如长的方法名定义为自己的类型
type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
