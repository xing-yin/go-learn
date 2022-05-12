package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/11
 * @Desc:
 */

type T7 struct {
	n int
	s string
}

func main() {
	var t = T7{
		n: 17,
		s: "hello,interface",
	}

	var ei interface{} = t // Go运行时使用 eface 结构表示ei
	fmt.Printf("%T\n", ei)
}
