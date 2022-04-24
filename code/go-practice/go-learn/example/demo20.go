package main

import (
	"fmt"
	"math"
)

type testInterface interface {
	test() int
}

type geometry2 interface {
	area() float64
	perim() float64
}

type rect2 struct {
	width, height float64
}

type circle2 struct {
	radius float64
}

func (r rect2) area() float64 {
	return r.width * r.height
}

func (r rect2) perim() float64 {
	return 2 * r.width * r.height
}

func (r rect2) test() int {
	return 1
}

func measure2(g geometry2) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func (c circle2) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle2) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle2) test() float64 {
	return 2
}

func main() {

	rect := rect2{
		width:  3,
		height: 4,
	}
	circle := circle2{radius: 5}

	measure2(rect)
	measure2(circle)

}
