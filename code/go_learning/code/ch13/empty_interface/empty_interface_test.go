package empty_interface

import (
	"fmt"
	"testing"
)

// 空接口类型
func DoSomething(p interface{}) {
	// 通过断⾔言来将空接口转换为指定类型
	//if i, ok := p.(int); ok {
	//	fmt.Println("Integer", i)
	//	return
	//}
	//if s, ok := p.(string); ok {
	//	fmt.Println("stirng", s)
	//	return
	//}
	//fmt.Println("Unknow Type")
	// 上面的简化写法
	switch v := p.(type) {
	case int:
		// 默认会有 break
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
	DoSomething(int64(20))
}
