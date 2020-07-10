package main

import "fmt"

func main() {
	var isSuccess bool
	isSuccess = true
	fmt.Println(isSuccess)

	//var a int = 8
	//var b int32 = 2
	//var c float64
	var d complex64 = 5 + 5i
	fmt.Println(d)

	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)

	s3 := "hello"
	s3 = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s3)

	s4 := `hello
		world
		string`
	fmt.Printf("%s\n", s4)

}
