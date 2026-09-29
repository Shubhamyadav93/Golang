package employee

import "fmt"

type Employee struct{
	Name string
	ID int 
	Department string
	Salary float64
}

// Unexported bonus calculation

func calculateBonus(salary float64) float64 {
	return salary * 0.10 // 10 % bonus 
}

func (e Employee) CalculateSalary() float64 {
	return e.Salary + calculateBonus(e.Salary)
}

func DisplayEmployee(e Employee)  {
	fmt.Println("Employee ID:",e.ID)
	fmt.Println("Name:",e.Name)
	fmt.Println("Total Salary (with Bonus):",e.CalculateSalary())
}


