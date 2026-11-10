package main

import "fmt"

func discount(coffeePrice *float32, discountNum float32) {
	*coffeePrice = *coffeePrice - (*coffeePrice * discountNum)
}

func main() {
	var coffeePrice float32 = 5.0
	var discountNum float32 = 0.10
	fmt.Printf("Basic coffee price is: %.2f\n", coffeePrice)
	fmt.Printf("discount is: %.2f\n", discountNum)
	discount(&coffeePrice, discountNum)
	fmt.Printf("Total coffee price is: %.2f\n", coffeePrice)
}
