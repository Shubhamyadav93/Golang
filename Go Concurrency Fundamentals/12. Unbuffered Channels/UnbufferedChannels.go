package main

import "fmt"

func main()  {
	// Step 1: Unbuffered channel create kiya (Capacity = 0)

	ch := make(chan int)

	// Step 2: Anonymous background Gorutine start ki 

	go func()  {
		ch <- 100 // Step 3: Value 100 ko channel me bheja (Send)
	}()

	// Step 4: Main Gorutine ne channel se value receive ki 

	value := <- ch 

	// Step 5: Terminal par print kiya 
	fmt.Println(value)// 100

}



