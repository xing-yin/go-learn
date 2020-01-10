package type_test

import "testing"

// 别名
type MyInt64 int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64

	// b = a
	b = int64(a)

	var c MyInt64
	//c = b
	c = MyInt64(b)

	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// 不支持指针运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	// %T 获取类型
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	// string 是值类型，默认初始值是空字符串，不是 nil
	t.Log("*" + s + "*")
	t.Log(len(s))
}
