package lec17

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/2
 * @Desc:
 */

type T1 int
type T2 T1
type T3 string

type (
	s1 int
	s2 s1
	s3 string
)

func main() {

	var n1 T1
	var n2 T2 = 5
	n1 = T1(n2)
	fmt.Println(n1)

	// var s T4 = "hi"
	// n1 = T1(s)

}
