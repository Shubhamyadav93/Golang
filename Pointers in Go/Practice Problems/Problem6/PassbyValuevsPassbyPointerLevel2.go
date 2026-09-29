package main

import "fmt"


func updaeValue(num int)  {
	num = 50
}

func updatePointer(num *int)  {
	*num = 50
}

func main()  {
	var num int
	fmt.Scan(&num)

	updaeValue(num)
	fmt.Println("After updateValue:",num)

	updatePointer(&num)
	fmt.Println("After updatePointer:",num)
}