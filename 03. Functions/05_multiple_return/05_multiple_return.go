package main

import "fmt"

func processPayment(totalAmount float64, tips float64, amountPaid float64) (float64, float64) {
	totalAmountDue := totalAmount + tips
	change := amountPaid - totalAmount
	return totalAmountDue, change
}

func main() {
	totalAmountDue, change := processPayment(6.50, 2.00, 10.00)
	fmt.Println("----------------------")
	fmt.Println("Total amount with tips is", totalAmountDue)
	fmt.Println("Change is ", change)
}
