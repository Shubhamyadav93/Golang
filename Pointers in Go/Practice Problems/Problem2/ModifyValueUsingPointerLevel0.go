package main

import "fmt"

func main()  {
	var num int
	fmt.Scan(&num)

	p:=&num
	*p = 100 
	fmt.Println(num)
}

