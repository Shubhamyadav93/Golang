package cart

import (
	"fmt"
	"productapp/product"
)

type Cart struct{
	Items []product.Product
}

func (c *Cart) AddProduct(p product.Product) {
	c.Items = append(c.Items, p)
}


func (c Cart) CalculateTotal() float64 {
	total := 0.0 
	for _, item := range c.Items{
		total += item.Price 
	}
	return total
}


func (c Cart) DisplayCart() {
	fmt.Println("---- Shopping Cart ---")
	for _, item := range c.Items{
		product.DisplayProduct(item)
	}
	fmt.Printf("Total Amount: Rs.%.2f",c.CalculateTotal())

}



