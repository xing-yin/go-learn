package client

import (
	// 注意这个路径是从 src 之后的目录开始（相对路径）
	"main/go_learning/code/ch15/series"
	"testing"
)

// 测试访问包外代码
func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))
	t.Log(series.Square(5))
	// 小写的函数不能被包外代码访问
	// t.Log(series.square(1))
}
