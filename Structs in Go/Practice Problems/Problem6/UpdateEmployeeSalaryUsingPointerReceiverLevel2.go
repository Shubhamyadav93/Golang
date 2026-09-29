package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func (e *Employee) updateSalary(newSalary float64) {
	e.Salary = newSalary
}

func main() {
	employee := Employee{
		Name:   "Priya",
		Salary: 40000,
	}

	fmt.Println("Before Update:")
	fmt.Println("Name:",employee.Name)
	fmt.Println("Salary:",employee.Salary)
	employee.updateSalary(50000)
	fmt.Println("After Salary:")
	fmt.Println("Name:",employee.Name)
	fmt.Println("Salary:",employee.Salary)

}


// Before Update:
// Name: Priya
// Salary: 40000
// After Salary:
// Name: Priya
// Salary: 50000

