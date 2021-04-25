package once_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

// 单例模式（还可以使用init()方法）

type Singleton struct {
	data string
}

var singleInstance *Singleton
var once sync.Once

// sync.Once 的 do 方法只执行一次
func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%X\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
