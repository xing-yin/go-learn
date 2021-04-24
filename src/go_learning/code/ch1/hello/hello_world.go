package main

// 包，表明代码所在的模块(包)

// 引⼊入代码依赖
import (
	"fmt"
	"os"
)

// 功能实现
func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1])
	}
}
