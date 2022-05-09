package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/9
 * @Desc:
 */

func main() {
	var i interface{} = 15
	println(i)
	type T6 struct{}
	i = "hello,go"

	var t T6
	i = t
	i = &t

	v, ok := i.(T6)
	fmt.Printf("%T\n", v)
	println(ok)
}
