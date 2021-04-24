package main

import "fmt"

/**
通道关闭
关闭一个通道意味着不能再向这个通道发送值了。这个特性可以用来给这个通道的接收方传达工作已经完成的信息。
*/
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("receive job:", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job:", j)
	}
	// 关闭通道
	close(jobs)
	fmt.Println("sent all jobs")

	// 程序会一直阻塞到收到消息
	<-done
}
