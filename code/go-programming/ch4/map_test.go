package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestMap(t *testing.T) {
	//funcName()

	//funcName2()

	//funcName3()

	funcName4()
}

// equal 判断 2 个 map 是否相等
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func funcName4() {
	ages := map[string]int{
		"alice": 20,
		"bob":   22,
	}
	if age, ok := ages["bob2"]; !ok {
		fmt.Printf("error:%d", age)
	}
}

func funcName3() {
	var ages map[string]int
	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)
	//ages["cindy"] = 21 //panic: assignment to entry in nil map
}

func funcName2() {
	ages2 := map[string]int{
		"alice": 20,
		"bob":   22,
	}
	fmt.Println(ages2)
	//delete(ages2, "alice")
	//fmt.Println(ages2)
	ages2["bob"]++
	//_ = &ages2["bob"] // compile error: cannot take address of map element

	var names []string
	for name := range ages2 {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages2[name])
	}

}

func funcName() {
	ages := make(map[string]int)
	fmt.Println(ages)
}
