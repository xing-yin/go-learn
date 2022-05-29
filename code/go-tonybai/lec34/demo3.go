package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/15
 * @Desc:
 */

type signal2 struct {
}

var ready2 bool

func worker2(i int) {
	fmt.Printf("worker %d:is working...\n", i)
	time.Sleep(time.Second)
	fmt.Printf("worker %d: works done\n", i)
}

func spawnGroup2(f func(i int), num int, groupSignal *sync.Cond) <-chan signal2 {
	c := make(chan signal2)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			groupSignal.L.Lock()
			for !ready2 {
				groupSignal.Wait()
			}
			groupSignal.L.Unlock()
			fmt.Printf("worker %d: start to work....\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal2(struct{}{})
	}()
	return c
}

func main() {
	fmt.Println("start a group of workers...")
	groupSignal := sync.NewCond(&sync.Mutex{})
	c := spawnGroup2(worker2, 5, groupSignal)

	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")

	groupSignal.L.Lock()
	ready2 = true
	groupSignal.Broadcast() // todo 一定要加上 Broadcast，否则会报错
	groupSignal.L.Unlock()

	<-c
	fmt.Println("the group of workers work done!")
}
