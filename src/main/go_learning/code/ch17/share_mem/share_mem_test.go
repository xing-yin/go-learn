package share_mem

import (
	"sync"
	"testing"
	"time"
)

// 线程不安全：结果不对，counter 在不同的协程中计数，存在竞争
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)

}

// 线程安全：此处用 sync.Mutex
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	// todo 运行的速度超过协程的运行速度，注释掉后结果会不准确：这种方式不如 sync.WaitGroup
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)

}

// sync.WaitGroup
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			// 协程完成
			wg.Done()
		}()
	}
	// WaitGroup 会在此处会阻塞，直到任务完成
	wg.Wait()
	t.Logf("counter = %d", counter)

}
