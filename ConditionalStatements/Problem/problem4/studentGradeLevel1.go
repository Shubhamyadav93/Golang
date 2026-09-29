package main

import "fmt"

func main() {
	marks := 86

	if marks >= 90 {
		fmt.Println("Grade A")
	}else if marks >= 80{
		fmt.Println("Grade B")
	}else if marks >= 70{
		fmt.Println("Grade C")
	}else if marks > 60{
		fmt.Println("Grade C")
	}else{
		fmt.Println("Fail")
	}
}


