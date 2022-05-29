package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/13
 * @Desc:
 */

var ch chan int

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int, 6)

	go func() {
		ch1 <- 11
	}()
	a := <-ch1

	ch2 <- 12
	b := <-ch2
	fmt.Println(a)
	fmt.Println(b)
}
