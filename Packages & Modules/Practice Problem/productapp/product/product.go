package product

import "fmt"

type Product struct{
	ID int 
	Name string
	Price float64
}

func DisplayProduct(p Product)  {
	fmt.Printf("[%d] %s - Rs.%.2f",p.ID,p.Name,p.Price)
}

