package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/4/25
 * @Desc:
 */

func main() {

	messages := make(chan string, 2)

	messages <- "buffer"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
