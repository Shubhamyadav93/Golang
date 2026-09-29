package main

import "fmt"

func swap(a,b *int)  {
	temp := *a 
	*a = *b 
	*b = temp
}

func main()  {
	var x , y int 
	fmt.Scan(&x,&y)
	swap(&x,&y)
	fmt.Println(x,y)
}
