package main

import (
	"calculatorapp/calculator"
	"fmt"
	"log"
)


func main()  {
	a,b := 10.0,5.0 
	fmt.Println("Add:",calculator.Add(a,b))
	fmt.Println("Subtract:",calculator.Subtract(a ,b))
	fmt.Println("Multiply:",calculator.Multiply(a,b))
	// fmt.Println("Divide:",calculator.DivideWithValidation(a,b,nil)) // ❌ Error


	// 🟢 SAHI TAREEQA: Function se result aur error dono receive karein

	result,err := calculator.DivideWithValidation(a,b)
	if err != nil{
		log.Fatalln("error during division:",err)
		return
	}
	fmt.Println("Divide:",result)


}

// Add: 15
// Subtract: 5
// Multiply: 50
// Divide: 2




