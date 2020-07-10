package flow_function

import (
	"fmt"
	"testing"
)

/// todo=================================if

func TestIf(t *testing.T) {

	// Go的if还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了
	//if x := computedValue(); x > 10 {
	//	fmt.Print("x is greater than 10")
	//} else {
	//	fmt.Print("x is greater than 10")
	//}
}

/// todo=================================goto
// Go有goto语句——请明智地使用它。用goto跳转到必须在当前函数内定义的标签。

// 例如假设这样一个循环：
func myFunc() {
	i := 0
Here: //这行的第一个词，以冒号结束作为标签
	println(i)
	i++
	goto Here //跳转到Here去
}

/// todo=================================for

// for配合range可以用于读取slice和map的数据：
func TestFor(t *testing.T) {
	map0 := make(map[string]int)
	map0["a"] = 1
	map0["b"] = 2
	for k, v := range map0 {
		fmt.Println("map key :", k)
		fmt.Println("map value :", v)
	}
}

// 由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用_来丢弃不需要的返回值 例如
func TestFor2(t *testing.T) {
	map0 := make(map[string]int)
	map0["a"] = 1
	map0["b"] = 2
	for _, v := range map0 {
		fmt.Println("map value :", v)
	}
}

/// todo=================================switch

// Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch,
// 但是可以使用fallthrough强制执行后面的case代码。
func TestSwith(t *testing.T) {
	integer := 6
	switch integer {
	case 4:
		fmt.Println("<4")
		fallthrough
	case 5:
		fmt.Println("<5")
		fallthrough
	case 6:
		fmt.Println("<6")
		fallthrough
	case 7:
		fmt.Println("<7")
		fallthrough
	case 8:
		fmt.Println("<8")
		fallthrough
	default:
		fmt.Println("default")
	}
}
