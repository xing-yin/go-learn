package person

import (
	"fmt"
	"testing"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/20
 * @Desc:
 */

func TestOlder(t *testing.T) {

}

// 变量的命名规范
func variable(t *testing.T) {
	//// 表达函数的预期输出和实际输出
	//// 方式1
	//if c.expected != actual {
	//	t.Fatalf("Expecte User-Agent '%s' does not match '%s'", c.expected, actual)
	//}
	//
	//// 方式2
	//if got != want {
	//	t.Errorf("wrong number got:%s, want:%s", got, want)
	//}
}

func TestMain(m *testing.M) {
	fmt.Println("do some setup")
	m.Run()
	fmt.Println("do some cleanup")
}
