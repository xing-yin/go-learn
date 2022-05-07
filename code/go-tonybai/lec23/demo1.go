package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/6
 * @Desc: 演示 panicking 过程
 */

func foo() {
	println("call foo")
	bar()
	println("exit foo")
}

func bar() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover the panic")
		}
	}()

	println("call bar")
	panic("panic occurs in bar")
	zoo()
	println("exit bar")
}

//func bar() {
//	println("call bar")
//	panic("panic occurs in bar")
//	zoo()
//	println("exit bar")
//}

func zoo() {
	println("call zoo")
	println("exit zoo")
}

func main() {
	println("call main")
	foo()
	println("exit main")
}
