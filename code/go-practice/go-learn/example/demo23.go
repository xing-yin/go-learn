package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/4/25
 * @Desc:
 */

func main() {

	message := make(chan string)

	go func() { message <- "ping" }()

	msg := <-message
	fmt.Println(msg)

}
