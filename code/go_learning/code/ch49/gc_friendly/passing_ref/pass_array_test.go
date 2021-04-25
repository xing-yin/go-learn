package gc_friendly

import (
	"testing"
)

const NumOfElems = 1000

type Content struct {
	Detail [10000]int
}

// 传值的gc
func withValue(arr [NumOfElems]Content) int {
	//	fmt.Println(&arr[2])
	return 0
}

// 传引用的gc
func withReference(arr *[NumOfElems]Content) int {
	//b := *arr
	//	fmt.Println(&arr[2])
	return 0
}

func TestFn(t *testing.T) {
	var arr [NumOfElems]Content
	//fmt.Println(&arr[2])
	withValue(arr)
	withReference(&arr)
}

func BenchmarkPassingArrayWithValue(b *testing.B) {
	var arr [NumOfElems]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(arr)
	}
	b.StopTimer()
}

func BenchmarkPassingArrayWithRef(b *testing.B) {
	var arr [NumOfElems]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withReference(&arr)
	}
	b.StopTimer()
}
