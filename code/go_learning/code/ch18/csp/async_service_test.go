package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	// 代码逻辑操作
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

// 串行执行（同步返回）
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	// 方式1：同步的channel
	//retCh := make(chan string)
	// 方式2:buffer channel(更高效)
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		// 向 channel 中放数据
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

//
func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}
