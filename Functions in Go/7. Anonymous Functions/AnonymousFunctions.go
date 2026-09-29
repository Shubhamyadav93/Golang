package main

import "fmt"

// 👉 Bina kisi naam ke bane function ko Anonymous Function kehte hain.

// Anonymous Function you can used inside the main function 



func main()  {

// 	func() {
// 	fmt.Println("Hello")
// }()

// Assigning an Anonymous Function to a Variable 

greet:= func()  {
	fmt.Println("Hello")
}
greet()
	
}

