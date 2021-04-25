package object_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	// 测试1：尝试放置超出池大小的对象
	// if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	// 	t.Error(err)
	// }
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			// 测试2:不释放并且1秒钟超时
			//fmt.Printf("%T\n", v)

			// 测试3:正常使用
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}

	}

	fmt.Println("Done")
}
