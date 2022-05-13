package main

import (
	"fmt"
	"time"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/13
 * @Desc: 有 bug 的输出（修复见 demo12）
 */

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- 5
		close(ch1)
	}()

	go func() {
		time.Sleep(time.Second * 7)
		ch2 <- 7
		close(ch2)
	}()

	var ok1, ok2 bool
	for {
		select {
		case x := <-ch1:
			ok1 = true
			fmt.Println(x)
		case x := <-ch2:
			ok2 = true
			fmt.Println(x)
		}

		if ok1 && ok2 {
			break
		}
	}
	fmt.Println("program end")
}
