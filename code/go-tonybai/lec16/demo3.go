package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/1
 * @Desc:
 */

func foo(m map[string]int) {
	m["k1"] = 11
	m["k2"] = 12
}

func main() {
	m := map[string]int{
		"k1": 1,
		"k2": 2,
	}

	fmt.Println(m)
	foo(m)
	fmt.Println(m)
}
