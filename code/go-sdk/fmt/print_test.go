package fmt

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	fmt.Print("param1", "|param2", "|param3\n")
	fmt.Println("param1", "|param2", "|param3")
	fmt.Printf("格式化输出：%s, 占位符：%d", "aaa", 12)
}
