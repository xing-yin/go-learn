package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"Alan", 20})

	fmt.Println(person{name: "Bob", age: 21})

	fmt.Println(person{name: "Cindy"})

	fmt.Println(&person{"David", 22})

	s := person{
		name: "Eva",
		age:  23,
	}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 24
	fmt.Println(sp.age)
}
