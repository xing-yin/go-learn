package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/13
 * @Desc: 演示 - 无缓冲 channel 还被用来实现 1 对 n 的信号通知机制
 */

func worker2(i int) {
	fmt.Printf("worker %d:is working...\n", i)
	time.Sleep(time.Second)
	fmt.Printf("worker %d: works done\n", i)
}

type signal2 struct {
}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal2) <-chan signal2 {
	c := make(chan signal2)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-groupSignal
			fmt.Printf("worker %d:start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal2{}
	}()
	return c
}

func main() {
	fmt.Println("start a group of workers...")
	groupSignal := make(chan signal2)
	c := spawnGroup(worker2, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	close(groupSignal)
	<-c
	fmt.Println("the group of workers work done!")
}
