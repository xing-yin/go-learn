package main

import "time"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/13
 * @Desc:
 */

func worker3(c chan int) {
	select {
	case <-c:
		// do something
	case <-time.After(30 * time.Second):
		return
	}
}

func main() {

}
