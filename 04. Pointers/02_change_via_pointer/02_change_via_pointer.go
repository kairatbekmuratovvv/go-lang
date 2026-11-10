package main

import "fmt"

func main() {
	var coffeePrice = 4.5
	fmt.Println("Coffee price is: ", coffeePrice)
	fmt.Println("Memory address: ", &coffeePrice)

	coffeePrice = 5.0
	fmt.Println(coffeePrice)
	fmt.Println("Memory address for 5.0: ", &coffeePrice)

	pointerCoffeePrice := &coffeePrice
	*pointerCoffeePrice = 6.0
	fmt.Println(coffeePrice)
	fmt.Println("Memory address for pointer: ", &pointerCoffeePrice)

}
