package ch6

import (
	"fmt"
	"math"
	"testing"
)

type Point struct {
	X, Y float64
}

// function
func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

// method: same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func TestPoint(t *testing.T) {
	p := Point{
		X: 1,
		Y: 2,
	}
	q := Point{
		X: 4,
		Y: 6,
	}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
}

//------------------------------------------------------------------------------------
type Path []Point

var path Path

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func TestPath(t *testing.T) {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}

//------------------------------------------------------------------------------------
func (p *Point) ScaleBy(factor float64) {
	p.X = factor
	p.Y = factor
}

func TestScaleBy(t *testing.T) {
	r := &Point{
		X: 1,
		Y: 2,
	}
	r.ScaleBy(2)
	fmt.Println(*r)

	p := Point{
		X: 2,
		Y: 3,
	}
	(&p).ScaleBy(2)
	fmt.Println(p)
}
