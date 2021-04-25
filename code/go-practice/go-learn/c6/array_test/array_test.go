package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr1 [3]int
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 4, 5}
	arr1[0] = 1
	t.Log(arr1)
	t.Log(arr2)
	t.Log(arr3)
	t.Log(arr1[0])
	t.Log(len(arr3))
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4}
	// way1
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	// way2
	for _, elem := range arr {
		t.Log(elem)
	}

	// way3
	for index, elem := range arr {
		t.Log(index, elem)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arrSection1 := arr[3:5]
	arrSection2 := arr[1:]
	arrSection3 := arr[:3]
	t.Log(arrSection1)
	t.Log(arrSection2)
	t.Log(arrSection3)
}
