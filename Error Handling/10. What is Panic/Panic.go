package main

import "fmt"


func CheckAge(age int)  {
	if age < 0 {
		// Triggers a fatal runtime panic
		panic("age cannot be negative")
	}
	fmt.Println("Age is:",age)
}

func main()  {
	fmt.Println("Starting program execution...")
	// ❌ This call trigger a panic because input (-5) is invalid 
	CheckAge(-5)
	
	//// ⚠️ This line will NEVER execute because panic halts the program execution
	fmt.Println("Program completed successfully")

}



// Starting program execution...
// panic: age cannot be negative

// goroutine 1 [running]:
// main.CheckAge(0x1004b7918?)
//         /Users/shubham/Desktop/Golang/Error Handling/10. What is Panic/Panic.go:9 +0x8c
// main.main()
//         /Users/shubham/Desktop/Golang/Error Handling/10. What is Panic/Panic.go:17 +0x54
// exit status 2

