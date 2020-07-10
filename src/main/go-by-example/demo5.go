package main

import "fmt"

/**
for 循环
for 是 Go 中唯一的循环结构。这里有 for 循环的三个基本使用方式。
*/
func main() {

	// 最常用的方式，带单个循环条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// 经典的初始化/条件/后续形式 for 循环
	for j := 1; j < 4; j++ {
		fmt.Println(j)
	}

	// 不带条件的 for 循环(相当于 while 循环)将一直执行，直到在循环体内使用了 break 或者 return 来跳出循环
	for {
		fmt.Println("loop")
		break
	}

}
