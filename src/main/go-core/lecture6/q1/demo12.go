package main

import "fmt"

// 判断变量的类型
var container = []string{"a", "b", "c"}

func main() {
	container := map[int]string{0: "a", 1: "b", 2: "c"}

	// 方式1
	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)
	if !(ok1 || ok2) {
		fmt.Printf("Error: unsupported container type: %T\n", container)
		return
	}
	fmt.Printf("The element is: %T. The container type is:%T\n", container[1], container)

	// 方式2：switch
	elem, err := getElement(container)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("The element is: %T. The container type is:%T\n", elem, container)

}

func getElement(containerI interface{}) (elem string, err error) {
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("Unsupported container type:%T\n", containerI)
		return
	}
	return
}
