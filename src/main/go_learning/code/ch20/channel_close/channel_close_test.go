package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

// 生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			// 往 channel 放数据
			ch <- i
		}
		// 发送完数据，关闭通道
		close(ch)
		// panic: send on closed channel
		// 向关闭的 channel 发送数据，会导致 panic
		//ch <- 10
		// WaitGroup
		wg.Done()
	}()

}

// 消费者
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			for i := 0; i < 15; i++ {
				fmt.Println(<-ch)
			}

			//// v, ok <-ch; ok 为 bool 值，true 表示正常接受，false 表示通道关闭
			//if data, ok := <-ch; ok {
			//	fmt.Println(data)
			//} else {
			//	fmt.Println("channel is closed")
			//	break
			//}
		}
		wg.Done()
	}()

}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()

}
