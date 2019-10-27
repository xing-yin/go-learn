package slice_test

import "testing"

// 切⽚片声明
func TestSliceInit(t *testing.T) {
	// slice 与数组的差别在于，不需要指定长度
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	// 其中len个元素会被初始化为默认零值，未初始化元素不不可以访问
	// 比如下面的 s2[3]
	//t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

// 切片的增长
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

// 切⽚片共享存储结构：查看长度与容量的大小
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	// 长度为3，容量为9(从year 的第3个到最后一个元素)
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	// 由于共享空间，一个修改会影响其他
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	//// 切片只能和nil比较
	//if a == b {
	//	t.Log("equal")
	//}
	t.Log(a, b)
}
