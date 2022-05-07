package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/4
 * @Desc:
 */

func main() {
	var m = map[string]int{
		"tony": 21,
		"tom":  22,
		"jim":  23,
	}

	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "tony")
		}
		counter++
		fmt.Println(k, v)
	}

	fmt.Println("counter is ", counter)
}
