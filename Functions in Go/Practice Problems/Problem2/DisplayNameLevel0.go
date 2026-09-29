package main

import "fmt"

func displayNme(name string)  {
	fmt.Println("Hello,",name)
}

func main()  {
	var name string
	fmt.Scan(&name)
	displayNme(name)
}



