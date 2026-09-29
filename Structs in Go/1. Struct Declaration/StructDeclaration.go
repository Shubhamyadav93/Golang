package main

import "fmt"

type Student struct{
	Name string
	Age int 
	Marks float64
	Department string
}


func main()  {

	// 3. Assigning Values to Struct Fields 

	// var s1 Student 
	// s1.Name = "Arun"
	// s1.Age = 20 
	// s1.Marks = 85.5 
	// s1.Department = "CSE"
	// fmt.Println(s1) //{Arun 20 85.5 CSE} 

							// 4 & 5. Struct Literal & Positional Initialization 
							// Named Fields (Recommended) 

	// s1 := Student {
	// 	Name: "Arun",
	// 	Age : 20,
	// 	Marks: 85.5,
	// 	Department: "CSE",
	// }
	// fmt.Println(s1)//{Arun 20 85.5 CSE}


						// Positional Initialization

	s1 := Student{"Arun",20,85.0,"CSE"}
	fmt.Println(s1)//{Arun 20 85 CSE}


	// 6 & 7. Accessing & Updating Struct Values 

	fmt.Println(s1.Name)//Access 
	s1.Marks = 95 
	fmt.Println(s1.Marks)

	
} 




