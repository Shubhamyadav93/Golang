package main

import (
	"errors"
	"fmt"
	"log"
)

// Function that performs alidation and returns explicit error
func Divide(a,b float64) (float64,error) {
	if b == 0{
		// Creating a explicit error value 
		return  0 , errors.New("cannot divide by zero")
	}
	// Operation success: returning result and nil error 
	return  a / b , nil 
}


func main()  {
	// Case 1: Valid Division 
	res1 , err := Divide(10,2)

	if err != nil{
		log.Println("Error:",err)
	}else{
		fmt.Printf("10 / 2 = %.2f\n",res1)
	}


	// Case 2: Invalid Division (Zero Division Error)

	res2,err := Divide(10,0)
	if err != nil{
		fmt.Println("[HADLED ERROR]:",err)
	}else{
		fmt.Printf("10 / 0 = %.2f\n",res2)
	} 
}


