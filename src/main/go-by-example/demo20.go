package main

import (
	"fmt"
	"math"
)

/**
接口
接口是方法特征的命名集合
更多关于 Go 的接口的知识自行搜索
*/

// 这里是一个几何体的基本接口
type geometry interface {
	area() float64
	perim() float64
}

// 在我们的例子中，我们将让 recta 和 circle 实现这个接口
type recta struct {
	width, height float64
}

type circle struct {
	redius float64
}

// ⚠️todo 要在 Go 中实现一个接口，我们只需要实现接口中的所有方法。
// 这里我们让 rect 实现了 geometry 接口。
func (r recta) area() float64 {
	return r.width * r.height
}

func (r recta) perim() float64 {
	return 2 * (r.width + r.height)
}

// circle 的接口实现
func (c circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.redius
}

// 如果一个变量是接口类型，那么我们可以调用这个被命名的接口中的方法。
// 这里有一个一通用的 measure 函数，利用这个特性，它可以用在任何 geometry 上。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := recta{width: 3, height: 4}
	c := circle{redius: 5}

	measure(r)
	measure(c)
}
