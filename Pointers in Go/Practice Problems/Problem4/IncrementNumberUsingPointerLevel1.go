package main

import "fmt"

func increment(num *int)  {
	*num++
}

func main()  {
	var n int 
	fmt.Scan(&n)
	increment(&n)
	fmt.Println(n)
	
}
