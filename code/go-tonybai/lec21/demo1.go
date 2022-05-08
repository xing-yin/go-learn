package main

import (
	"fmt"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/5
 * @Desc:
 */

func myAppend(sl []int, elems ...int) []int {
	fmt.Printf("%T\n", elems)
	if len(elems) == 0 {
		println("no elems to append")
		return sl
	}

	sl = append(sl, elems...)
	return sl
}

func main() {
	sl := []int{1, 2, 3}
	sl = myAppend(sl) // no elems to append
	fmt.Println(sl)

	sl = myAppend(sl, 4, 5, 6)
	fmt.Println(sl)
}
