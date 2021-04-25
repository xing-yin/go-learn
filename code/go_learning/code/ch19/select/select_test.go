package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	//time.Sleep(time.Millisecond * 50)
	// 超时测试
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	//retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

// select 做超时机制
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	//	超时设置
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
