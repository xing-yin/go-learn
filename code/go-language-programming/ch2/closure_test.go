package ch2

import (
	"fmt"
	"testing"
)

func TestClosure(t *testing.T) {
	var j int = 5

	// 在变量a指向的闭包函数中，只有内部的匿名函数才能访问变量i，而无法通过其他途径访问到，因此保证了i的安全性。
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i,j: %d,%d\n", i, j)
		}
	}()

	// 变量a指向的闭包函数引用了局部变量i和j，i的值被隔离，在闭包外不 能被修改，改变j的值以后，再次调用a，发现结果是修改过的值。
	a()
	j *= 2
	a()
}
