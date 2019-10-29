package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

// 协程示例
func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 关键字 go 开启协程：i 是值传递，利用拷贝，不存在竞争关系
		go func(i int) {
			//time.Sleep(time.Second * 1)
			fmt.Println(i)
		}(i)

		// ⚠️错误写法：下面这种方式i 被共享了，需要加锁，有竞争关系
		//go func() {
		//	fmt.Println(i)
		//}()
	}
	// 程序运行很快，可以加个等待
	time.Sleep(time.Millisecond * 50)
}
