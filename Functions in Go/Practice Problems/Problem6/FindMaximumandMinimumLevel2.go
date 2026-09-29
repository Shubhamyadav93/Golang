package main

import "fmt"

// func findMinMax(a, b int) (int, int) {
// 	return max(a, b), min(a, b)
// }

// OR 

func findMinMax(a,b int) (max int, min int) {
	if a >= b{
		return a, b
	}
	return b,a
}


func main() {
	var a, b int
	fmt.Scan(&a,&b)

	Maximum , Minimum := findMinMax(a,b)
	fmt.Println("Maximum = ", Maximum)
	fmt.Println("Minimum =",Minimum)

}

