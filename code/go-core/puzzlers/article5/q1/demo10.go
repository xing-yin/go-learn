package main

import "fmt"

// 重名变量===>"强龙不压地头蛇"
var block = "package"

func main() {
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
}
