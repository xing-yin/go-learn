package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name           string
	LifeExpectance int
}

func (b Bird) Fly() {
	fmt.Println("I am flying")
}

func main() {
	sparrow := &Bird{
		Name:           "Sparrow",
		LifeExpectance: 3,
	}
	s := reflect.ValueOf(sparrow).Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, f.String(), f.Type(), f.Interface())
	}
}
