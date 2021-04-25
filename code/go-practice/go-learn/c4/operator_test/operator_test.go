package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	// "=="数组比较，比较的不是引用，而是数组元素
	t.Log(a == b)
	// 数组维度不同，无法比较
	//t.Log(a == c)
	t.Log(a == d)
}

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	t.Log(Readable, Writeable, Executable)
	t.Log(a&^Readable, a&^Writeable, a&^Executable)
}
