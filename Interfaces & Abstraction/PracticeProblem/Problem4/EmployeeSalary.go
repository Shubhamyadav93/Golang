package main

import "fmt"

type Employee interface{
	CalaculateSalary() float64
	GetName() string
}

type FullTimeEmployee struct{
	Name string
	MonthlySalary float64
}

func (f FullTimeEmployee) CalaculateSalary() float64 {
	return f.MonthlySalary
}

func (f FullTimeEmployee) GetName() string {
	return f.Name 
}

type PartTimeEmployee struct{
	Name string
	HourseWorked float64
	HourlyRate float64 
}

func (p PartTimeEmployee) CalaculateSalary() float64{
	return p.HourseWorked * p.HourlyRate
}

func (p PartTimeEmployee) GetName() string {
	return p.Name
}


func displaySalary(e Employee)  {
	fmt.Printf("%s Salary: %.0f \n", e.GetName()  ,e.CalaculateSalary())
}



func main()  {
	f := FullTimeEmployee{
		Name: "Rahul",
		MonthlySalary: 50000,
	}
	p:= PartTimeEmployee{
		Name: "Priya", 
		HourseWorked: 80,
		HourlyRate: 500,
	}

	displaySalary(f)
	displaySalary(p)
}

// Rahul Salary: 50000 
// Priya Salary: 40000 


