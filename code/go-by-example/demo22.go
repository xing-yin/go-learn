package main

import "fmt"

/**
协程
Go 协程 在执行上来说是轻量级的线程
*/

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 假设有一个函数 f(s)。我们使用一般的方式调并同时运行。
	f("direct")

	// 使用 go f(s) 在一个 Go 协程中调用这个函数。
	// 这个新的 Go 协程将会并行的执行这个函数调用，异步运行
	go f("goroutine")

	// 也可以为匿名函数启动一个 Go 协程，异步运行
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 现在这两个 Go 协程在独立的 Go 协程中异步的运行，所以我们需要等它们执行结束。这里的 Scanln 代码需要我们在程序退出前按下任意键结束
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

	// 当我们运行这个程序时，将首先看到阻塞式调用的输出，然后是两个 Go 协程的交替输出。
	// 这种交替的情况表示 Go 运行时是以 异步 的方式运行协程的
}
