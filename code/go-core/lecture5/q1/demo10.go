package main

import "fmt"

// 变量重名 和 变量的重声明
var block = "package"

func main() {
	block := "function"
	{
		block := "inner"
		fmt.Printf("the block is %s.\n", block)
	}
	fmt.Printf("the blcok is %s.\n", block)
}
