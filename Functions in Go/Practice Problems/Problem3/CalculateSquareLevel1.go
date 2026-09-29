package main

import "fmt"

func square(n int) int{
	return n * n
}

func main()  {
	var n int
	fmt.Scan(&n)
	fmt.Println("Square =", square(n))	// 16 
}


