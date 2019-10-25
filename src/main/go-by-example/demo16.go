package main

import "fmt"

/**
函数-递归
Go 支持 递归。这里是一个经典的阶乘示例
*/
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	res := fact(7)
	fmt.Println(res)
}
