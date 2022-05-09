package main

/**
 * @Author: Alan Yin
 * @Date: 2022/5/7
 * @Desc:
 */

//
//type MyInt *int
//// r的基类型为指针类型，编译器报错：invalid receiver type MyInt (MyInt is a pointer type)
//func (r MyInt) String() string {
//	return fmt.Sprintf("%d", *(*int)(r))
//}
//
//type MyReader io.Reader
//// r的基类型为接口类型(MyReader)，编译器报错：invalid receiver type MyReader (MyReader is an interface type) return r.Read(p)}
//func (r MyReader) Read(p []byte) (int, error) {
//	return r.Read(p)
//}

func main() {

}
