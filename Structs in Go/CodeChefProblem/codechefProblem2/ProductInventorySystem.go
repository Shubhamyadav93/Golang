package main

import "fmt"

type Product struct{
	ID string
	Name string
	Price float64
	InStock bool
}

func CreateProduct(id,name string,price float64,inStock bool)Product{
	return Product{
		ID : id,
		Name: name,
		Price: price,
		InStock: inStock,
	}

}

func FormatSummary(p Product)string{
	return fmt.Sprintf("ID: %s | Name:%s | Price : %2.f | InStock: %t",p.ID,p.Name,p.Price,p.InStock)
}

func main()  {
	p:=CreateProduct("101","Yadhu",29.99,true)
	fmt.Println(FormatSummary(p))
}


// ID: 101 | Name:Yadhu | Price : 30 | InStock: true