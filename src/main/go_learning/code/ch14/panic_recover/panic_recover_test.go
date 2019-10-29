package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

// panic vs. os.Exit
func TestPanicVxExit(t *testing.T) {

	// recover:恢复
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))

	//os.Exit(-1)
	//fmt.Println("End")
}
