package main

import (
	"fmt"
	"os"
	"strings"
)

/**
环境变量
环境变量是一个在为 Unix 程序传递配置信息的普遍方式。
让我们来看看如何设置，获取并列举环境变量。
*/
func main() {

	// 使用 os.Setenv 来设置一个键值队。使用 os.Getenv获取一个键对应的值。如果键不存在，将会返回一个空字符串。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// 使用 os.Environ 来列出所有环境变量键值队。
	// 这个函数会返回一个 KEY=value 形式的字符串切片。
	// 你可以使用strings.Split 来得到键和值。这里我们打印所有的键。
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], pair[1])
	}

}

// 如果我们在运行前设置了 BAR 的值，那么运行程序将会获取到这个值。
//$ BAR=2 go run demo61.go
//FOO: 1
//BAR: 2
//...
