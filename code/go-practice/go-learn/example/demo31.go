package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/4/25
 * @Desc:
 */

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
	fmt.Println("finish")
}
