package main

/**
 * @Author: Alan Yin
 * @Date: 2022/5/9
 * @Desc: Go 接口类型允许嵌入的不同接口类型的方法集合存在交集，但前提是交集中的方法名字和函数签名部分要保持一致
 */

type Interface1 interface {
	M1()
}

type Interface2 interface {
	M1(string)
	M2()
}

type Interface3 interface {
	Interface1
	//Interface2 // Duplicate method 'M1'
}
