package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 多返回值的函数
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// todo 可以复用（函数式编程） 运行时间时长计算： 输入是 inner func(op int) int 这是其他函数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		// 这是真正运行的函数，有点想装饰函数
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

// 耗时函数，用于验证 timeSpent 函数
func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

// 可变参数（本质是一个数组）
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	// go 中类似抛出异常
	panic("err")
	fmt.Println("won't excute")
}
