package main

import (
	"foodapp/food"
	"foodapp/order"
)

func main()  {
	f1 := food.Item{
		Name: "Burger",
		Price: 250,
	}
	f2 := food.Item{
		Name: "Fries",
		Price: 120,
	}

	myOrder := order.Order{
		ID: 101,

	}

	myOrder.AddItem(f1)
	myOrder.AddItem(f2)
	myOrder.DisplayOrder()
}