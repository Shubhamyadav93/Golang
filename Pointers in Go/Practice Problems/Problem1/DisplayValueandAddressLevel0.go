package main

import "fmt"


func main()  {
	var num int 
	num = 10 
	p:=&num 
	fmt.Println("Value:",num)
	fmt.Println("Address:",p)
	fmt.Println("Pointer Value:",*p)

}

