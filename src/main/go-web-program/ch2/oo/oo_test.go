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
// type typeName typeLiteral

type age int
type money float32
type months map[string]int

func Test2(t *testing.T) {
	m := months{
		"January":  31,
		"February": 28,
		// ...
	}
	fmt.Println(m)
}

// 实际上只是一个定义了一个别名,有点类似于c中的typedef，例如上面ages替代了int

// 03 可以在任何的自定义类型中定义任意多的method
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box //a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (b1 BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range b1 {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (b1 BoxList) PaintItBlack() {
	for i, _ := range b1 {
		b1[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func Test3(t *testing.T) {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())

}

// 上面的代码通过const定义了一些常量，然后定义了一些自定义类型:
// - Color作为byte的别名
// - 定义了一个struct:Box，含有三个长宽高字段和一个颜色属性
// - 定义了一个slice:BoxList，含有Box

//然后以上面的自定义类型为接收者定义了一些method:
// - Volume()定义了接收者为Box，返回Box的容量
// - SetColor(c Color)，把Box的颜色改为c
// - BiggestColor()定在在BoxList上面，返回list里面容量最大的颜色
// - PaintItBlack()把BoxList里面所有Box的颜色全部变成黑色
// - String()定义在Color上面，返回Color的具体颜色(字符串格式)

/// todo=================================指针作为receiver

/// todo=================================method 继承
// Go 有一个神奇之处，method也是可以继承的：
// 如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。

// 01 让我们来看下面这个例子
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名字段
	school string
}

type Employee struct {
	Human
	company string
}

// 在 Human 上定义一个 method
func (h *Human) SayHi() {
	fmt.Printf("Hi，I am %s you can call me on %s\n", h.name, h.phone)
}

func TestHuman(t *testing.T) {
	mark := Student{Human{"Mark", 25, "6666"}, "MIT"}
	sam := Employee{Human{"Sam", 24, "77777"}, "google"}

	mark.SayHi()
	sam.SayHi()
}

/// todo=================================method 重写
// Human 定义 method
func (h *Human) SayHi2() {
	fmt.Printf("Hi，I am %s you can call me on %s\n", h.name, h.phone)
}

// Student 的 method 重写 Human 的 method
func (s *Student) SayHi2() {
	fmt.Printf("Hi，I am %s you can call me on %s, my school is %s\n", s.name, s.phone, s.school)
}

// Employee 的 method 重写 Human 的 method
func (e *Employee) SayHi2() {
	fmt.Printf("Hi，I am %s you can call me on %s, my company is %s\n", e.name, e.phone, e.company)
}

func TestHuman2(t *testing.T) {
	mark := Student{Human{"Mark", 25, "6666"}, "MIT"}
	sam := Employee{Human{"Sam", 24, "77777"}, "google"}

	mark.SayHi2()
	sam.SayHi2()
}

// 通过这些内容，我们可以设计基本的面向对象的程序了，但是 Go 里面的面向对象是如此的简单，没有任何的私有、公有关键字，
// 通过大小写来实现(大写开头的为公有，小写开头的为私有)，方法也同样适用这个原则。
