package main

import (
	"productapp/cart"
	"productapp/product"
)

func main()  {

	p1 := product.Product{
		ID: 1 ,
		Name: "Laptop",
		Price: 50000,
	}
	p2 := product.Product{
		ID: 2,
		Name: "Mouse",
		Price: 1000,
	}

	myCart := cart.Cart{}
	myCart.AddProduct(p1)
	myCart.AddProduct(p2)
	myCart.DisplayCart()
}

