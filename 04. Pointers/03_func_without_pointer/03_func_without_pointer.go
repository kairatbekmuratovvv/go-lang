package main

import "fmt"

func discount(coffeePrice float32, discount float32) float32 {
	discountPrice := coffeePrice - (coffeePrice * discount)
	return discountPrice
}
func main() {
	var coffeePrice float32 = 5.0
	fmt.Printf("Basic coffee price is: %.2f\n", coffeePrice)
	coffeePriceWithDiscount := discount(coffeePrice, 0.10)
	fmt.Printf("Coffee price with discount: %.2f\n", coffeePriceWithDiscount)

}
