package ch6

import (
	"fmt"
	"testing"
	"time"
)

func TestMethodValue(t *testing.T) {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance // method value
	fmt.Println(distanceFromP(q))

	var origin Point
	fmt.Println(distanceFromP(origin))
}

//------------------------------------------------------------------------------------
type Rocket struct {
}

func (r Rocket) Launch() {
}

func TestRocket(t *testing.T) {
	r := new(Rocket)
	time.AfterFunc(10*time.Second, func() { r.Launch() })
	time.AfterFunc(10*time.Second, r.Launch) //用方法"值"更加简洁
}

//------------------------------------------------------------------------------------
func TestMethodExpression(t *testing.T) {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance // method expression
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}
