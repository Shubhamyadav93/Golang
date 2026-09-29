package students

import "fmt"

// Struct Definition (Outside function)

type Students struct{
	Name string
	Age int 
}

// Function Definition (Outside struct)

func Display(s Students)  {
	fmt.Println("Name:",s.Name)
	fmt.Println("Age:",s.Age)
}


