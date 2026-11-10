package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Espresso":   2.50,
		"Americano":  3.00,
		"Latte":      3.50,
		"Cappuccino": 0.00,
	}

	drink := "Cappuccino"
	fmt.Printf("Cappuccino price: $%.2f\n", menu[drink])

	price, exists := menu[drink]
	fmt.Println("Price:", price)
	fmt.Println("Exists", exists)

	if exists {
		fmt.Printf("Price: $%.2f\n", price)
	} else {
		fmt.Println(drink, "is not on the menu")
	}

}
