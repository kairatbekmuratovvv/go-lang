package main

import "fmt"

func main() {
	taxRate := 0.10

	calculateTax := func(amount float64) float64 {
		return amount * taxRate
	}
	subtotal := 25.00
	tax := calculateTax(subtotal)
	total := subtotal + tax
	fmt.Printf("Total is %.2f\n", total)
}
