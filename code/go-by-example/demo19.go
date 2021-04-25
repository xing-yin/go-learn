package main

import "fmt"

/**
方法
Go 支持在结构体类型中定义方法
*/

// 长方形
type rect struct {
	width, height int
}

// 可以为值类型或者指针类型的接收器定义方法。
// 这里的 area 方法有一个指针接收器类型 *rect
func (r *rect) area() int {
	return r.width * r.height
}

// 这里是一个值类型接收器的例子
func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{width: 10, height: 5}

	// 这里我们调用上面为结构体定义的两个方法
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go 自动处理方法调用时的值和指针之间的转化。
	// 你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接受的数据
	rp := &r
	rp.width = 12
	rp.height = 6
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
