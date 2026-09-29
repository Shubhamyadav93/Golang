package main

import "fmt"

func main()  {
	Math := 82 
	Physics := 78 
	Chemistry := 84 

	average := (Math+Physics+Chemistry)/3
	if Math >= 75 && Physics >= 70 && Chemistry >= 70 && average >= 80{
		fmt.Println("Eligible for Admission")
	}else{
		fmt.Println("Not Eligible")
	}

}


