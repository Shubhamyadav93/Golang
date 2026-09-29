package main

import "fmt"

// 1. Dependency Rule (Contract)

type Ingredient interface{
	GetItems() string
}

// 2. Concrete Ingredients (Bahar se supply honge)

type FreshVegetables struct{}

func (f FreshVegetables) GetItems()string {
	return "Fresh organic Tomatoes & Paneer"
}

// 3. Chef (Component jo dependency receive karega)

type Chef struct{
	ingredient Ingredient // 👈 Dependency injected from outside!
}

func (c Chef) Cook() {
	items := c.ingredient.GetItems()
	fmt.Println("Chef is cooking dish using:",items)
}

func main()  {
	   // Step 1: Supplier ne ingredients tayar kiye (Outside)
	mySupply := FreshVegetables{}

	   // Step 2: Chef ko ingredients supply/inject kiye gaye

	chef := Chef{
		ingredient: mySupply,
	}

	   // Step 3: Chef ne bas cooking ki!

	   chef.Cook()

}

// Chef is cooking dish using: Fresh organic Tomatoes & Paneer 


