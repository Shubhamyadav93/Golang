package main

import "fmt"

func arithmeticops(a,b int) (int,int,int,int){
	return  a + b , a - b , a * b , a / b 
}

func main()  {
	var a , b int
	fmt.Scan(&a,&b)

	sum,diff,prod,quot := arithmeticops(a,b)
	fmt.Println("Sum=",sum)
	fmt.Println("Difference=",diff)
	fmt.Println("prduct =",prod)
	fmt.Println("Quotient =",quot)

}



