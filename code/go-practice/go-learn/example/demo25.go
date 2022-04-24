package main

import (
	"fmt"
	"time"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/4/25
 * @Desc:
 */

func worker(done chan bool) {
	fmt.Println("working....")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	<-done
}
