package main

import (
	"fmt"
	"time"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/8
 * @Desc:
 */

type field struct {
	name string
}

//// 错误的写法
//func (p *field) print() {
//	fmt.Println(p.name)
//}

// 正确的写法
func (p field) print() {
	fmt.Println(p.name)
}

func main() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print()
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go v.print()
	}

	time.Sleep(3 * time.Second)
}
