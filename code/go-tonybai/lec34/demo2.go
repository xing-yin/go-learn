package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/15
 * @Desc: 演示-用sync.Mutex 实现对条件轮询等待
 */

type signal struct {
}

var ready bool

func worker(i int) {
	fmt.Printf("worker %d:is working...\n", i)
	time.Sleep(time.Second)
	fmt.Printf("worker %d: works done\n", i)
}

func spawnGroup(f func(i int), num int, mu *sync.Mutex) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				mu.Lock()
				if !ready {
					mu.Unlock()
					time.Sleep(100 * time.Microsecond)
					continue
				}
				mu.Unlock()
				fmt.Printf("worker %d: start to work....\n", i)
				f(i)
				wg.Done()
				return
			}
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal(struct{}{})
	}()
	return c
}

func main() {
	fmt.Println("start a group of workers...")
	mu := &sync.Mutex{}
	c := spawnGroup(worker, 5, mu)

	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")

	mu.Lock()
	ready = true
	mu.Unlock()

	<-c
	fmt.Println("the group of workers work done!")
}
