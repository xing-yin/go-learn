package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/12
 * @Desc:
 */

func main() {

	go fmt.Println("I am a goroutine")

	var c = make(chan int)
	go func(a, b int) {
		c <- a + b
	}(3, 4)

}
