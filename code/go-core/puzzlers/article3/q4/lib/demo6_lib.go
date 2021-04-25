package lib

import (
	in "go-learn/code/go-core/puzzlers/article3/q4/lib/internal"
	"os"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
