package test

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type OrderByAge []Person

func (a OrderByAge) Len() int {
	return len(OrderByAge{})
}

func (a OrderByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a OrderByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func ExamplePerson() {
	people := []Person{
		{
			Name: "Allen",
			Age:  20,
		},
		{
			Name: "Bob",
			Age:  22,
		},
		{
			Name: "Cindy",
			Age:  21,
		},
	}

	fmt.Println(people)
	sort.Sort(OrderByAge(people))
	fmt.Println(people)

	// Output:
	//[Allen: 20 Bob: 22 Cindy: 21]
	//[Allen: 20 Bob: 22 Cindy: 21]
}
