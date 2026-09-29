package order

import (
	"fmt"
	"foodapp/food"
)

type Order struct{
	ID int 
	Items []food.Item
}

// Pointer Receiver to mutate (add items to) the Order struct 

func (o *Order) AddItem(item food.Item) {
	o.Items = append(o.Items, item)
}

// Value Receiver to compute total and display summary 


func (o Order) DisplayOrder() {
	fmt.Printf("Order #%d Summary:\n",o.ID)
	total := 0.0 
	for _,item := range o.Items{
		fmt.Printf("- %s: Rs.%.2f\n",item.Name,item.Price)
		total += item.Price
	}
	fmt.Printf("Total Payable: Rs.%.2f\n",total)

}




