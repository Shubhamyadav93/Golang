package main

import "fmt"

type Address struct{
	City string
	State string
}

type Student struct{
	Name string
	Age int 
	Address // BINA FIELD NAME KE RAKHA (Embedded Field) 
}

func main()  {
	s1:= Student{
		Name: "Arun",
		Age: 20,
		Address: Address{
			City: "Delhi",
			State: "Delhi",
		},
	}

	// DIRECT ACCESS (Field Promotion): s1.Address.City ki jagah direct s1.City  

	fmt.Println("City:",s1.City)  // Output: Delhi 

}


