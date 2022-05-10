package employee

/**
 * @Author: Alan Yin
 * @Date: 2022/5/9
 * @Desc: 演示：接口类型的类型嵌入
 */

type E interface {
	M1()
	M2()
}

//type I interface {
//	M1()
//	M2()
//	M3()
//}

// I 与上面等价
type I interface {
	E
	M3()
}

func main() {

}
