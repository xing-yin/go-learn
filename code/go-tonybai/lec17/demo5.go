package lec17

import (
	"bytes"
	"fmt"
	"sync"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/2
 * @Desc:
 */

//type T1 struct {
//	t2 T2
//}
//
//type T2 struct {
//	t1 T1
//}

type T4 struct {
	t  *T4
	st []T31
	m  map[string]T31
}

func main() {
	var mu sync.Mutex
	mu.Lock()
	// do something
	mu.Unlock()

	var b bytes.Buffer
	b.Write([]byte("hello, Go"))
	fmt.Println(b.String())
}
