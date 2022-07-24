package ch5

import (
	"os"
	"testing"
)

func TestRiskLoop(t *testing.T) {
	filenames := []string{}
	badSample(filenames)

	betterSample(filenames)
}

func betterSample(filenames []string) error {
	for _, filename := range filenames {
		if err := doFile(filename); err != nil {
			return err
		}
	}
	return nil
}

func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

// 在循环体中的defer语句需要特别注意，因为只有在函数执行完毕后，这些被延迟的函数才会执行。下面的代码会导致系统的文件描述符耗尽，因为在所有文件都被处理之前，没有文件会被关闭。
func badSample(filenames []string) error {
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close() // NOTE: risky; could run out of file
	}
	return nil
}
