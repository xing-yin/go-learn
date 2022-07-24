package ch5

import (
	"fmt"
	"os"
	"testing"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func TestSum(t *testing.T) {
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3, 4))

	vals := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(vals...))
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func TestErrorf(t *testing.T) {
	linenum, name := 12, "count"
	errorf(linenum, "undedined:%s", name)
}
