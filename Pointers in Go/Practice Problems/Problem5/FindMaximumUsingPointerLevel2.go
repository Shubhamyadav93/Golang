package main

import "fmt"



func maximum(a,b *int)  {
	if *a >= *b{
		fmt.Println(*a)
	}else{
		fmt.Println(*b)

	}

}

func main(){
	var x , y int 
	fmt.Scan(&x,&y)
	maximum(&x,&y)
}
