package main

import "fmt"

/**
老牌入门案例: hello world
*/
func main() {
	fmt.Println("hello world")
}

/**
方式1： 要运行这个程序，将这些代码放到 demo2.go 中并且使用 go run 命令。
> go run httpdemo.go

方式2:
有时候我们想将程序编译成二进制文件。可以通过 go build 命来达到目的。
> go build httpdemo.go

然后我们可以直接运行这个二进制文件。
> ./demo1
*/
