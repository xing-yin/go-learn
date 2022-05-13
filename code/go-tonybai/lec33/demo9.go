package main

import (
	"log"
	"sync"
	"time"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/13
 * @Desc: 演示-带缓冲 channel 用作计数信号量的例子
 */

var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func main() {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- (i + 1)
		}
		close(jobs)
	}()

	var wg sync.WaitGroup

	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			active <- struct{}{}
			log.Printf("handle job")
			time.Sleep(3 * time.Second)
			<-active
			wg.Done()
		}(j)
	}

	wg.Wait()
}
