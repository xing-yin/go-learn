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
> go run demo1.go

方式2:
有时候我们想将我们的程序编译成二进制文件。我们可以通过 go build 命来达到目的。
> go build demo1.go

然后我们可以直接运行这个二进制文件。
> ./demo1
*/
