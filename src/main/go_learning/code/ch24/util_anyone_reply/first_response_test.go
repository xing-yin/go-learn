package concurrency

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/**
仅需任意任务完成
*/

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	// 如果不指定buffer channel，不会阻塞【会导致协程泄漏】
	//ch := make(chan string)
	// todo 可以防止协程泄漏
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	// 一旦 channel 中有数据就可以返回
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	// runtime.NumGoroutine() 输出当前协程数
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())

}
