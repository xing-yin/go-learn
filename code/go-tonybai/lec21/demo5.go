package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/5
 * @Desc: 利用闭包简化函数调用
 */

func times(x, y int) int {
	return x * y
}

// partialTimes 函数返回一个函数func(int)，这个函数返回一个int
func partialTimes(x int) func(int) int {
	return func(y int) int {
		return times(x, y)
	}
}

func main() {
	times(2, 5)
	times(3, 5)
	times(4, 5)

	timesTwo := partialTimes(2) // 以高频乘数 2 为固定乘数的乘法函数
	timesThree := partialTimes(3)
	timesFour := partialTimes(4)

	fmt.Println(timesTwo(5)) // 等价于 times(2,5)
	fmt.Println(timesTwo(6))
	fmt.Println(timesThree(5))
	fmt.Println(timesThree(6))
	fmt.Println(timesFour(5))
	fmt.Println(timesFour(6))
}
