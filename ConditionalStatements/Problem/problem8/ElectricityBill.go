package main

import "fmt"

func main()  {
	units := 250 

	bill := 0 
	if units <= 100{
		bill = units * 2
	}else if units <= 200{
		bill = (100 * 2) + ((units - 100)*5)
	}else{
		bill = (100 * 2) + (100 * 5) + ((units - 200) * 8)
	}

	fmt.Println(bill)

}

