package type_test

import "testing"

// 别名
type MyInt int64

// 类型转化
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	// 1.Go 语⾔言不不允许隐式类型转换
	//a = b
	//b = a
	// 支持显示类型转换
	b = int64(a)

	var c MyInt
	// 2.别名和原有类型也不不能进⾏行行隐式类型转换
	//c = b
	c = MyInt(b)
	t.Log(a, b, c)
}

// 指针类型
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// 1. 不不⽀支持指针运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	// %T 获得类型
	t.Logf("%T %T", a, aPtr)
}

// 字符串默认初始值
func TestString(t *testing.T) {
	var s string
	// 2. string 是值类型，其默认的初始化值为空字符串串，⽽而不不是 nil
	// 初始化零值是“”，不需要对 string == nil 的判断
	t.Log("*" + s + "*")
	t.Log(len(s))

}
