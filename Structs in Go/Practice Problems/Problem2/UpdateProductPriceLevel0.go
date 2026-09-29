package main

import "fmt"

type Product struct{
	ID int 
	Name string
	Price float64
}

func main()  {
	product := Product{
		ID: 101,
		Name: "Laptop",
		Price: 50000,
	}

	fmt.Println("Before Update:")
	fmt.Println("ID:", product.ID)
	fmt.Println("Name:",product.Name)
	fmt.Println("Price:",product.Price)

	product.Price =45000

	fmt.Println("After Update:")
	fmt.Println("ID:",product.ID)
	fmt.Println("Name:", product.Name)
	fmt.Println("Price:",product.Price)

}

// Before Update:
// ID: 101
// Name: Laptop
// Price: 50000
// After Update:
// ID: 101
// Name: Laptop
// Price: 45000

