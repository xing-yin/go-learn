package lec13

import "fmt"

/**
 * @Author: Alan Yin
 * @Date: 2022/4/30
 * @Desc:
 */

//const (
//	Apple, Banana = 11, 22
//	Strawberry, Grape
//	Peer, Watermelon
//)

const (
	Apple, Banana = iota, iota + 10
	Strawberry, Grape
	Peer, Watermelon
)

const (
	mutexLocked1 = 1 << iota
	mutext2
	mutext3
	mutext4 = iota
	mutext5 = 1e5
)

const (
	_ = iota
	one
	two
	three
)

const (
	_ = iota
	Pin1
	Pin2
	Pin3
	_
	Pin5
)

func main() {

	//fmt.Println(Apple)
	//fmt.Println(Banana)
	//fmt.Println(Strawberry)
	//fmt.Println(Grape)
	//fmt.Println(Peer)
	//fmt.Println(Watermelon)

	//fmt.Println("---------------")
	//fmt.Println(mutexLocked1)
	//fmt.Println(mutext2)
	//fmt.Println(mutext3)
	//fmt.Println(mutext4)
	//fmt.Println(mutext5)

	//fmt.Println("---------------")
	//fmt.Println(one)
	//fmt.Println(two)
	//fmt.Println(three)

	fmt.Println("---------------")
	fmt.Println(Pin1)
	fmt.Println(Pin2)
	fmt.Println(Pin3)
	fmt.Println(Pin5)
}
