package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

// 接口实现
type GoProgrammer struct {
}

// 实现方法: go 的 duck type 设计
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
