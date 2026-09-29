package main

import "fmt"

// 👉 range:-  keyword ka use collections (Arrays, Slices, Strings, Maps, Channels) par iterate (ek-ek element scan) karne ke liye hota hai

// for index, value := range collection {
// // statements
// }

func main()  {
	//Example with Slice
	// nums:= []int{10,20,30,40}
	// for index, value := range nums{
	// 	fmt.Println(index,value)
	// }


	// Ignoring the Index

	nums := []int{10,2,30}
	for _, value := range nums{
		fmt.Println(value)
	}

	// Range with String 

	str := "Go"
	for index,value:= range str{
		fmt.Println(index , value)
	}
	// 0 71
	// 1 111		
	

}


