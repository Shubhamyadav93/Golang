package main

import "fmt"

func main() {
	a := 15
	b := 28
	c := 11

	if a >= b && a >= c {
		fmt.Println(a)
	}else if b >= a && b >= c{
		fmt.Println(b)
	}else{
		fmt.Println(c)
	}
	
}



