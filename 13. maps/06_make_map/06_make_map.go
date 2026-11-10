package main

import "fmt"

func main() {
	//var stock map[string]int
	//fmt.Println(stock)
	//stock = make(map[string]int)
	stock := make(map[string]int)

	stock["Espresso"] = 10
	stock["Latte"] = 25

	fmt.Println("Products in stock:", stock)

	for drink, count := range stock {
		fmt.Println(drink, count)
	}
}
