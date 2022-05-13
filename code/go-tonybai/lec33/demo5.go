package main

import (
	"fmt"
	"time"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/13
 * @Desc: 演示 1 对 1 通知信号
 */

type signal struct {
}

func worker() {
	println("worker is working...")
	time.Sleep(1 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		c <- signal{}
	}()
	return c
}

func main() {
	println("start a worker")
	c := spawn(worker)
	<-c
	fmt.Println("worker work done")
}
