package main

import "fmt"

func main() {
	coffeeTypes := []string{"Espresso", "Cappuccino", "Latte", "Americano"}
	fmt.Println("Original menu:", coffeeTypes)
	fmt.Println("Len of:", len(coffeeTypes), "Capacity is:", cap(coffeeTypes))

	for _, coffee := range coffeeTypes {
		print(" ", coffee, "\n")
	}

}
