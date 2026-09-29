package main

import "fmt"

func calculateSalary(basic, hraPct, daPct float64) (hra float64, da float64, gross float64) {
	hra = basic * (hraPct / 100)
	da = basic * (daPct / 100)
	gross = basic + hra + da
	return
}

func main() {
	var basic, hraPct, daPct float64
	fmt.Scan(&basic,&hraPct,&daPct)
	hra,da,gross := calculateSalary(basic,hraPct,daPct)

	fmt.Println("HRA = %Of",hra)
	fmt.Println("DA = %Of",da)
	fmt.Println("Gross Salary = %Of",gross)

}


