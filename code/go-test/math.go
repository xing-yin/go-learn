package test

import (
	"math"
	"math/rand"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/20
 * @Desc: 演示如何编写测试用例
 */

// Abs returns the absolute value of x.
func Abs(x float64) float64 {
	return math.Abs(x)
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Min(x, y float64) float64 {
	return math.Min(x, y)
}

func RandInt() int {
	return rand.Int()
}
