package main

import "fmt"


											// Nested Structs (Struct Inside Struct) 
// 1. Inner Struct (Chhota Struct)
type Address struct{
	city string
	State string
}

// 2. Outer Struct (Bada Struct jiske andar Inner Struct hai)

type Student struct{
	Name string
	Age int 
	Addr Address // Nested Struct Field 
}

func main()  {
	// Object Initialization (Tarika 1: Nested Style)
	s1 := Student{
		Name: "Arun",
		Age: 20,
		Addr: Address{
			city: "Bengaluru",
			State: "Karnataka",
		},
	}

	// Accessing Values (Chaining Dot Operator)
	fmt.Println("Name:",s1.Name) 
	fmt.Println("City:",s1.Addr.city) // Output: Bengaluru
	fmt.Println("State:",s1.Addr.State) // Output: Karnataka
	
	
	// Updating Value 
	s1.Addr.city = "Mumbai"
	fmt.Println("Update City:",s1.Addr.city) // Output: Mumbai 


}

// Name: Arun
// City: Bengaluru
// State: Karnataka
// Update City: Mumbai


