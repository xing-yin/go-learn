package main

import "fmt"

/**
老牌入门案例: hello world
*/
//func main() {
//	//fmt.Println("hello world")
//	fmt.Println(fmt.Sprintf("\"error\":\"server_error\"\n \"message\": %+v", "msg"))
//}

/**
方式1： 要运行这个程序，将这些代码放到 demo2.go 中并且使用 go run 命令。
> go run httpdemo.go

方式2:
有时候我们想将程序编译成二进制文件。可以通过 go build 命来达到目的。
> go build httpdemo.go

然后我们可以直接运行这个二进制文件。
> ./demo1
*/

// 看恢复后，整个程序的执行情况是什么样的// 解释 panicking 过程：
func foo() {
	println("call foo")
	bar()
	println("exit foo")
}

// 不同的函数
func bar() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover the panic:", e)
		}
	}()

	println("call bar")
	panic("panic occurs in bar")
	zoo()
	println("exit bar")
}

func zoo() {
	println("call zoo")
	println("exit zoo")
}

func main() {
	println("call main")
	foo()
	println("exit main")
}
