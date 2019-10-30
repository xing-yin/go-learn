package concat_string

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/**
高效的字符串串连接
*/

const numbers = 100

// 性能不好
func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
	b.StopTimer()
}

// 字符串连接【最推荐】
func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var builder strings.Builder
		for i := 0; i < numbers; i++ {
			builder.WriteString(strconv.Itoa(i))

		}
		_ = builder.String()
	}
	b.StopTimer()
}

// 字符串连接推荐
func BenchmarkBytesBuf(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var buf bytes.Buffer
		for i := 0; i < numbers; i++ {
			buf.WriteString(strconv.Itoa(i))
		}
		_ = buf.String()
	}
	b.StopTimer()
}

// 性能不好
func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s += strconv.Itoa(i)
		}

	}
	b.StopTimer()
}
