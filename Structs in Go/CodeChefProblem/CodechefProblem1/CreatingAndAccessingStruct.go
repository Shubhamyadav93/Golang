package main

import "fmt"

type User struct {
	ID   string
	Name string
	Age  int
}

func main() {
	// Create Struct instance
	u := User{
		ID:   "101",
		Name: "Yadhu",
		Age:  19,
	}
	// Access and Print struct field
	fmt.Printf("User: %s (Id: %s, Age: %d)",u.Name,u.ID,u.Age)// // User: Yadhu (Id: 101, Age: 19)%                                                                                                                                                              
	fmt.Println(u.Name,u.ID,u.Age)//Yadhu 101 19

}
