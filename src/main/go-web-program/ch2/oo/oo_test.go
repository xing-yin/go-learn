package oo

import (
	"fmt"
	"math"
	"testing"
)

/// todo=================================method
// method的语法如下：
// func (r ReceiverType) funcName(parameters) (results)

// 01
type Rectangle struct {
	length, width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func TestMethod(t *testing.T) {
	r := Rectangle{10, 5}
	c := Circle{3}
	fmt.Println("Area of r is:", r.area())
	fmt.Println("Area of c is:", c.area())
}

// ⚠️在使用method的时候重要注意几点:
// - 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
// - method里面可以访问接收者的字段
// - 调用method通过.访问，就像struct里面访问字段一样

// 值得说明的一点是，图示中method用虚线标出，意思是此处方法的Receiver是以值传递，而非引用传递，Receiver还可以是指针,
// 两者的差别在于: 指针作为Receiver会对实例对象的内容发生操作,而普通类型作为Receiver仅仅是以副本作为操作对象,并不对原实例对象发生操作。

// 02
// 是不是method只能作用在struct上面呢？
// 当然不是咯，他可以定义在任何你自定义的类型、内置类型、struct等各种类型上面。
// 什么叫自定义类型，自定义类型不就是struct嘛，不是这样的哦，struct只是自定义类型里面一种比较特殊的类型而已，还有其他自定义类型申明，可以通过如下这样的申明来实现。
