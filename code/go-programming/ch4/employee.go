package main

import "time"

type Employee struct {
	ID        int
	Name      string
	Address   string
	Dob       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func foo() {
	position := &dilbert.Position
	*position = "Senior " + *position

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (team player)"
}

func GetEmployeeByID(id int) *Employee {
	return &Employee{}
}
