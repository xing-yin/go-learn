package array_slice

import "testing"

func TestSum(t *testing.T) {

	nums := [5]int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 15

	if want != got {
		t.Errorf("got %d want %d given %v", got, want, nums)
	}
}
