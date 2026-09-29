package main

import "fmt"

func main()  {
	var income float64 = 85000

	var tax float64 
	if income <= 300000{
		tax = 0 
	}else if income <= 700000{
		tax = income * 0.05 
	}else if income <= 1000000{
		tax = income * 0.10
	}else{
		tax = income * 0.20
	}

	fmt.Println("% Of",tax)
}

