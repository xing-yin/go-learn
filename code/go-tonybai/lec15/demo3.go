package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/1
 * @Desc:
 */

func main() {
	//var s []int
	//fmt.Println(s)
	//s = append(s, 1)
	//fmt.Println(len(s), cap(s))
	//s = append(s, 2)
	//fmt.Println(len(s), cap(s))
	//s = append(s, 3)
	//fmt.Println(len(s), cap(s))
	//s = append(s, 4)
	//fmt.Println(len(s), cap(s))
	//s = append(s, 5)
	//fmt.Println(len(s), cap(s))

	u := [...]int{1, 2, 3, 4, 5}
	fmt.Println("arr:", u)

	s := u[1:3]
	fmt.Printf("Slice(len=%d,cap=%d):%v\n", len(s), cap(s), s)

	s = append(s, 11)
	fmt.Println("after append 11,arr:", u)
	fmt.Printf("after append 11,Slice(len=%d,cap=%d):%v\n", len(s), cap(s), s)

	s = append(s, 12)
	fmt.Println("after append 12,arr:", u)
	fmt.Printf("after append 12,Slice(len=%d,cap=%d):%v\n", len(s), cap(s), s)

	s = append(s, 13)
	fmt.Println("after append 13,arr:", u)
	fmt.Printf("after append 13,Slice(len=%d,cap=%d):%v\n", len(s), cap(s), s)
}
