package string_test

import (
	"strconv"
	"strings"
	"testing"
)

// 常⽤用字符串串函数

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	// 字符串分割
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	// 字符串连接
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	// 字符串与其他类型的转换(如整数)
	s := strconv.Itoa(10)
	t.Log("str" + s)
	// 因为有2个返回值，不能直接这样操作
	//t.Log(10 + strconv.Atoi("10"))
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
