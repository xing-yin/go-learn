package ch3

import (
	"fmt"
	"testing"
)

func TestInt(t *testing.T) {
	//intOver()

	//printBitTest()

	//typeConvert()

	lostPrecision()
}

// 精度丢失
func lostPrecision() {
	f := 3.141
	i := int(f)
	fmt.Println(f, i)

	f = 1.99
	fmt.Println(int(f))
}

func typeConvert() {
	var a int32 = 1
	var b int16 = 2
	var c int = int(a) + int(b)
	fmt.Println(c)
}

func printBitTest() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}

func intOver() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)
}
