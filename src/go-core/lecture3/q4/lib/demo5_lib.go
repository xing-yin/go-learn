package lib

// todo  强制：为了不让使用代码包的使用者感到困惑，我们总是应该让声明的包名与其父目录的名称保持一致

import (
	in "main/go-core/lecture3/q4/internal"
	"os"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
