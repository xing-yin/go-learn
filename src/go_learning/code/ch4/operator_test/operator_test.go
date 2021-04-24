package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

// 用 == ⽐较数组
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	// 维度不相同不能比较
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

// &^ 按位置零
func TestBitClear(t *testing.T) {
	a := 7 //0111
	a = a &^ Readable
	a = a &^ Executable
	t.Log(Readable)
	t.Log(Writable)
	t.Log(Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
