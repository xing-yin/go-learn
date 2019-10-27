package constant_test

import "testing"

// 常量量定义
// 快速设置连续值:会自动递增加1（简洁写法）
const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

// 连续位常量
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1 //0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
