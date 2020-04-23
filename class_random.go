package main

import (
	"fmt"
)

func minMax(a, b int) (min, max int) {
	if a > b {
		min = b
		max = a
	} else {
		max = b
		min = a
	}
	return min, max
}

func testdef() {
	fmt.Println("last execuation")
}

func testPointer(a *int) {
	*a = *a + 2
}

// Student ...
type Student struct {
	roll int
	name string
}

func (s Student) getStudentName() string {
	return s.name
}

// Salary is ...
type Salary struct {
	basic, hr, ta float64
}

// Employee is ...
type Employee struct {
	firstName     string
	lastName      string
	email         string
	age           int
	monthlySalary []Salary
}

func (e Employee) getEmployeeInfo() string {
	fmt.Println(e.firstName, e.lastName, e.email)
	fmt.Println(e.age)
	for _, info := range e.monthlySalary {
		fmt.Println("--------------------------")
		fmt.Println(info.basic)
		fmt.Println(info.hr)
		fmt.Println(info.ta)
	}
	return "_______________________"
}

func main() {
	e := Employee{
		firstName: "shohag",
		lastName:  "mia",
		email:     "shohag.fks@gmail.com",
		age:       34,
		monthlySalary: []Salary{
			Salary{
				basic: 34,
				hr:    4,
				ta:    22,
			},
			Salary{
				basic: 345,
				hr:    33,
				ta:    3,
			},
		},
	}
	fmt.Println(e.getEmployeeInfo())
}
