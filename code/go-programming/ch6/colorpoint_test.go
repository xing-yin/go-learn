package ch6

import (
	"fmt"
	"image/color"
	"testing"
)

type ColorPoint struct {
	Point // 嵌入结构体
	Color color.RGBA
}

func TestColorPoint(t *testing.T) {
	var cp ColorPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"
}

type ColorPoint2 struct {
	*Point // 嵌入结构体
	Color  color.RGBA
}
