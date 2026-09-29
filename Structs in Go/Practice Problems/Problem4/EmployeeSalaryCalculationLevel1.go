package main

import "fmt"

type Employee struct{
	ID int 
	Name string
	BasicSalary float64
}

func calculateSalary(e Employee) float64  {
	hra := e.BasicSalary * 0.20 
	da := e.BasicSalary * 0.10 
	return e.BasicSalary + hra + da 

}

func main()  {
	emp := Employee{
		ID: 101,
		Name: "Rahul",
		BasicSalary: 30000,
	}

	gross := calculateSalary(emp)
	fmt.Println("Employee:", emp.Name)
	fmt.Println("Gross Salary: %.Of",gross) 

}

// Employee: Rahul
// Gross Salary: %.Of 39000

