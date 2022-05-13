package main

import (
	"fmt"
	"sync"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/13
 * @Desc:
 */

type counter2 struct {
	c chan int
	i int
}

func NewCounter() *counter2 {
	cter := &counter2{
		c: make(chan int),
	}
	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	return cter
}

func (cter *counter2) Increase2() int {
	return <-cter.c
}

func main() {
	cter := NewCounter()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cter.Increase2()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
