package main

import "io"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/9
 * @Desc: 初识接口
 */

type MyInterface interface {
	M1(int) error
	M2(io.Writer, ...string)
}

//// 等价的写法
//type MyInterface interface {
//	M1(a int) error
//	M2(writer io.Writer, strs ...string)
//}

func main() {

}
