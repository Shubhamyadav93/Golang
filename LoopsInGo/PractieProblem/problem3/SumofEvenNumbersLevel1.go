package main

import "fmt"

func main()  {
	var n int = 10 

	sum:=0
	for i:=1;i<=n;i++{
		if i%2 == 0{
			sum += i 
		}
	}
	fmt.Println("Sum =",sum)

}



