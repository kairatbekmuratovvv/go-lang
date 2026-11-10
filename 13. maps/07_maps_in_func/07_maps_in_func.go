package main

import "fmt"

func sellCoffee(stock map[string]int, coffeeType string, quantity int) {
	currentStockQty, exists := stock[coffeeType]

	if exists {
		if currentStockQty >= quantity {
			stock[coffeeType] -= quantity
			fmt.Printf("Sold %d %s(s)\n", quantity, coffeeType)
			fmt.Println("Stock in the function:", stock)
		} else {
			fmt.Printf("Not enough %d in stock. Only %d left.", quantity, currentStockQty)
		}
	} else {
		fmt.Printf("%s is not available in stock.\n", coffeeType)
	}

}

func main() {
	coffeeStock := map[string]int{
		"Espresso":   10,
		"Latte":      5,
		"Cappuccino": 8,
	}

	sellCoffee(coffeeStock, "Matcha", 1) //Not available
	sellCoffee(coffeeStock, "Espresso", 5)
	sellCoffee(coffeeStock, "Espresso", 1)
	fmt.Println("Stock in the function:", coffeeStock["Espresso"])
	sellCoffee(coffeeStock, "Latte", 7) // not enough in stock

}
