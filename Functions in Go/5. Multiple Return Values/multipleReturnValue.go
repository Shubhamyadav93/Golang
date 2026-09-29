package main

import "fmt"

func calculate(a int, b int) (int ,int){
	add := a+b 
	sub := a - b 
	return add, sub
}


func main()  {
	sum,difference := calculate(10,5)
	fmt.Println(sum)
	fmt.Println(difference)
}
