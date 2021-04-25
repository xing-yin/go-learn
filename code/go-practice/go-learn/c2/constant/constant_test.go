package constant

import "testing"

// 常量定义
const a = 1

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

// 连续位常量
const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTry2(t *testing.T) {
	t.Log(Readable, Writeable, Executable)
}
