package tempcobv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

var c Celsius
var f Fahrenheit

func testCompare() {
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//fmt.Println(c == f) // Invalid operation: c == f (mismatched types Celsius and Fahrenheit)
	fmt.Println(c == Celsius(f))
}

func (c Celsius) String() {
	fmt.Sprintf("%g C", c)
}
