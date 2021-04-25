package main

import (
	"fmt"
	"math"
)

func main() {
	r := rectangle{
		weight: 10,
		height: 5,
	}
	c := circular{radius: 4}

	measure(r)
	measure(c)
}

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	weight, height float64
}

func (r rectangle) area() float64 {
	return r.weight * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.weight)
}

type circular struct {
	radius float64
}

func (c circular) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circular) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}
