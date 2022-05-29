package main

import (
	"fmt"
	"math"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/26
 * @Desc:
 */

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
