package fmt

import (
	"fmt"
	"testing"
)

func TestFormatCommon(t *testing.T) {
	user := &User{Id: 2}
	fmt.Printf("%v\n", user)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%#v\n", user)
	fmt.Printf("%T\n", user)
	fmt.Printf("%%\n")
}

func TestFormatBool(t *testing.T) {
	fmt.Printf("%t\n", true)
	fmt.Printf("%t\n", false)
}

func TestFormatInt(t *testing.T) {
	n := 69
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	fmt.Printf("%U\n", n)
	a := 96
	fmt.Printf("%q\n", a)
	fmt.Printf("%q\n", 0x4E2D)
}

func TestFormatFloat(t *testing.T) {
	f := 18.54
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%F\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)
}

func TestFormatStringAndByteArray(t *testing.T) {
	s := "这是字符串"
	b := []byte{65, 66, 67}
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", b)
	fmt.Printf("%q\n", s)
	fmt.Printf("%q\n", b)
	fmt.Printf("%x\n", s)
	fmt.Printf("%x\n", b)
	fmt.Printf("%X\n", s)
	fmt.Printf("%X\n", b)
}

func TestFormatPoint(t *testing.T) {
	u := &User{Id: 1}
	fmt.Printf("%p\n", u)
}

func TestFormatWidth(t *testing.T) {
	n := 13.14
	fmt.Printf("%f\n", n)
	fmt.Printf("%10f\n", n)
	fmt.Printf("%10s\n", "这是字符串")
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%10.2f\n", n)
	fmt.Printf("%10.f\n", n)
}

func TestFormatOtherFlag(t *testing.T) {
	s := "这是字符串"
	fmt.Printf("%    d\n", 10)
	fmt.Printf("%    d\n", -10)
	fmt.Printf("%s\n", s)
	fmt.Printf("%10s\n", s)
	fmt.Printf("%-10s\n", s)
	fmt.Printf("%10.2f\n", 13.14)
	fmt.Printf("%-10.2f\n", 13.14)
	// 前面使用 0 补齐
	fmt.Printf("%010s\n", s)
}
