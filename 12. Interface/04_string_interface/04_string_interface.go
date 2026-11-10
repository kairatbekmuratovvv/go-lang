package main

import "fmt"

type Order struct {
	Customer string
	Item     string
	Quantity int
}

func (o Order) String() string {
	return fmt.Sprintf("Order: %s has ordered %s (%d)", o.Customer, o.Item, o.Quantity)
}

func main() {
	orderOne := Order{
		Customer: "Kiro",
		Item:     "Americano",
		Quantity: 1,
	}
	fmt.Println(orderOne)

}
