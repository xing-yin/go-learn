package remote

import (
	"testing"
	// 通过 go get 来获取远程依赖
	// cm 是自定义别名
	cm "github.com/easierway/concurrent_map"
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
