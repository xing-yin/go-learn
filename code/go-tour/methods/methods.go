package main

import (
	"fmt"
	"math"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/26
 * @Desc: Methods (https://go.dev/tour/methods/1)
 */

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
