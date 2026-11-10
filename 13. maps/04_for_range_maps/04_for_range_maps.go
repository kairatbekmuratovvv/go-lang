package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Espresso":   2.50,
		"Americano":  3.00,
		"Latte":      3.50,
		"Cappuccino": 3.25,
	}
	for drink, price := range menu {
		fmt.Println(drink, price)
	}
}
