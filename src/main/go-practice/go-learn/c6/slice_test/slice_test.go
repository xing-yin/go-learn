package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	// slice 相比于数组，不需要指定长度
	//var s0 []int
	//t.Log(len(s0),cap(s0))
	//s0 = append(s0,1,2,3)
	//t.Log(len(s0),cap(s0))

	//s1:= []int {1,2,3}
	//t.Log(len(s1),cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	// 未初始化的切片元素不可访问
	//t.Log(s2[0],s2[1],s2[2],s2[3])
	s2 = append(s2, 2)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))

}

// 切片的增长(容量以2为次幂)
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}
