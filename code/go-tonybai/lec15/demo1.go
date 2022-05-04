package main

import (
	"fmt"
	"unsafe"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/1
 * @Desc:
 */

func foo(arr [5]int) {
}

func main() {

	var arr1 [5]int
	//var arr2 [6]int
	//var arr3 [5]string

	foo(arr1)
	//foo(arr2)
	//foo(arr3)

	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("arr len is:", len(arr))
	fmt.Println("arr size is:", unsafe.Sizeof(arr))

	var arr4 = [...]int{21, 22, 23}
	fmt.Printf("%T\n", arr4)

	var arr5 = [...]int{99: 79}
	fmt.Printf("%T\n", arr5)
	fmt.Println(arr5)
	fmt.Println(arr5[0], arr5[18])
	//fmt.Println(arr5[-1])
	//fmt.Println(arr5[1000])

}
