package ch2

import (
	"fmt"
	"testing"
)

func modify(arr [5]int) {
	arr[0] = 66
	fmt.Println("In modify(),arr values:", arr)
}

func TestModify(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	modify(arr)
	fmt.Println("In main(),arr values:", arr)
}
