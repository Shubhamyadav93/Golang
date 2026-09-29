package main

import "fmt"


func main()  {
	var n int = 100 

	found := false 
	for i := 1; i<= n; i++{
		if i%7 == 0 && i%11==0{
			fmt.Println("First Number =",i)
			found = true
			break
		}
	}

	if !found{
		fmt.Println("Not Found")
	}
}
