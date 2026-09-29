package main

import "fmt"


func main()  {
	x := 10 
	p := &x 
	*p = 50//ka matlab: p me rakhe address par jao aur wahan ki value badal kar 50 kar do.
	fmt.Println(x)// 50 
}
