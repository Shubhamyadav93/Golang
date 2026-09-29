package main

import "fmt"


func calculateSI(p, r, t float64) float64 {
	return (p * r * t /100)
}

func main()  {
	var p,r,t float64 
	fmt.Scan(&p,&r,&t)
	si := calculateSI(p,r,t)
	fmt.Println("Simple Interest =", si)
}

