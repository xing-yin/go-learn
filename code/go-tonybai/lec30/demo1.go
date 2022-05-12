package main

/**
 * @Author: Alan Yin
 * @Date: 2022/5/11
 * @Desc:
 */

func Add1(a int64, b int64) int64 {
	return a + b
}

// Adder 过度设计的嫌疑
type Adder interface {
	Add(int64, int64) int64
}

func Add(adder Adder, a int64, b int64) int64 {
	return adder.Add(a, b)
}

func main() {

}
