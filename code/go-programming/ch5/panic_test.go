package ch5

import "fmt"

func Parse(input string) error {
	defer func() {
		if p := recover(); p != nil {
			err := fmt.Errorf("internal error: %v", p)
			fmt.Println(err)
		}
	}()
	return nil
}
