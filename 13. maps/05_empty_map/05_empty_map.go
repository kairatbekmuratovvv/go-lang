package main

import "fmt"

func main() {
	var stock map[string]int
	fmt.Println(stock)
	fmt.Printf("Location in memory: %p\n", &stock)

	if stock == nil {
		fmt.Println("Stock is nil!")
	}

	//stock["Espresso"] = 10
	//stock["Latte"] = 25
	//
	//fmt.Println("Products in stock:", stock)
}
