package main

import "fmt"

func main()  {
	var n int = 10

	for i:=1;i<=n;i++{
		if i%3==0{
			continue
		}
		fmt.Println(i)
	}
}

