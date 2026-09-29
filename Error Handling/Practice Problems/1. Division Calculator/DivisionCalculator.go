package main

import (
	"errors"
	"fmt"
)


func Divide(a,b int) (int,error) {
	if b == 0{
		// Division by zero is mathematically undefined -> Return zero value & error
		return 0, errors.New("cannot divide by zero")
	}
	// Happy Path -> Return calculate quotient and nil
	return a / b , nil
}

func main()  {
	// Attempting division with b = 0 
	res,err :=Divide(10,0)

	// Guard Clause Pattern 
	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	fmt.Println("Result:",res)
}





