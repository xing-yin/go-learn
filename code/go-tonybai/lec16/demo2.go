package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/1
 * @Desc:
 */

func main() {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}

	fmt.Printf("{")
	for k, v := range m {
		fmt.Printf("[%d:%d]", k, v)
	}
	fmt.Printf("}")

	fmt.Println("")
	for k := range m {
		fmt.Printf("[%d]", k)
	}

	fmt.Println("")
	for _, v := range m {
		fmt.Printf("[%d]", v)
	}
}
