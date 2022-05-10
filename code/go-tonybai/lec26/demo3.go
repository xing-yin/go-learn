package employee

import (
	"fmt"
	"reflect"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/9
 * @Desc:
 */

type I5 interface {
	M1()
	M2()
}

type T5 struct {
	I5
}

func (T5) M3() {
}

func dumpMethodSet(i interface{}) {
	dynType := reflect.TypeOf(i)

	if dynType == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynType.NumMethod()
	if n == 0 {
		fmt.Printf("%s 's method set is empty!\n", dynType)
		return
	}

	fmt.Printf("%s 's method set :\n", dynType)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynType.Method(j).Name)
	}
	fmt.Printf("\n")
}

func main() {
	var t T5
	var p *T5
	dumpMethodSet(t)
	dumpMethodSet(p)
}
