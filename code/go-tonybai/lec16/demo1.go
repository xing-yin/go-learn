package main

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/1
 * @Desc:
 */

type Position struct {
	x float64
	y float64
}

func main() {
	//a := map[string]int{}
	////var b map[string]int
	//fmt.Println(a == nil)
	//a["k1"] = 1
	//
	//m2 := map[Position]string{
	//	Position{1.1, 2.1}: "school",
	//	Position{1.2, 2.2}: "school2",
	//}
	//fmt.Println(m2)
	//
	//m3 := map[Position]string{
	//	{1.1, 2.1}: "school",
	//	{1.2, 2.2}: "school2",
	//}
	//fmt.Println(m3)

	//m4 := make(map[int]string)
	//m5 := make(map[int]string, 8)
	//fmt.Println(m4)
	//fmt.Println(m5)

	m6 := make(map[int]string)
	m6[1] = "v1"
	m6[2] = "v2"
	m6[3] = "v3"
	m6[1] = "v111"
	fmt.Println(m6)
	delete(m6, 3)
	fmt.Println(m6)

	//	fmt.Println(len(m6))
	//	m6[4] = "v4"
	//	fmt.Println(len(m6))

	//m7 := make(map[string]int)
	//v, ok := m7["key1"]
	//if !ok {
	//	// "key1"不在 map 中
	//} else {
	//	// key1 在 map 中, v 被赋值为 key1 键对应的 value
	//	fmt.Println(v)
	//}

}
