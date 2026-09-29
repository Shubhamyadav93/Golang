package main

import "fmt"

func main()  {
	age:= 25 
	slaray:= 45000
	creditScore := 760

	if age >= 21 && slaray >= 30000 && creditScore >= 700{
		fmt.Println("Eligible")
	}else{
		fmt.Println("Not Eligible")
	}
}


