package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/4/25
 * @Desc:
 */

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 0; j < 4; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}

	close(jobs)
	fmt.Println("send all jobs")

	<-done
}
