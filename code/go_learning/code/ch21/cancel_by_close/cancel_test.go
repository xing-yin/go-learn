package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// 检查取消的channel上是否有取消的消息
func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

// 只有有消息即可取消（空的可以取消）【不推荐】
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

// 利用关闭广播也可以关闭
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	//cancel_1(cancelChan)
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}
