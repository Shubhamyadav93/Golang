package main

import "fmt"

// Value Semantics (Pass by Value) 

// func update(x int)  {
// 	x = 100 
// }

// func main()  {
// 	num := 10
// 	update(num)
// 	fmt.Println(num)// 10 
// }


// Pointer Semantics (Pass by Pointer/Reference) 


func update(x *int)  {
	*x = 100
}


func main()  {
	num := 10
	update(&num)
	fmt.Println(num)

}


