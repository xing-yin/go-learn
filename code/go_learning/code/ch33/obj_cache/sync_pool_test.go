package object_pool

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		// 不存在时会创建
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	//GC 会清除sync.pool中缓存的对象
	//runtime.GC()
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

// 测试协程下的缓存机制
func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}

	// 产生了 3个后面只需要再创建7个
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
