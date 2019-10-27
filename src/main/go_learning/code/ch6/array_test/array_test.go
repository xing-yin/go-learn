package array_test

import "testing"

// 数组的声明
func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	// 会自动匹配长度
	arr3 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
	t.Log(len(arr3))
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	// 方式1
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	// 方式2:下划线 _ 可以忽略它
	for _, elem := range arr3 {
		t.Log(elem)
	}
	for idx, elem := range arr3 {
		t.Log(elem)
		t.Log(idx)
	}
}

// 数组截取
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3Section := arr3[2:len(arr3)]
	t.Log(arr3Section)
}
