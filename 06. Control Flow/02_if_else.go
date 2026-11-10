package main

import "fmt"

func main() {
	customerName := "Kiro"
	orderSize := "Large"

	if orderSize == "Small" {
		fmt.Println(customerName, "ordered a Small coffee. It will be ready in 2 min!")
	} else {
		fmt.Println(customerName, "ordered something bigger. It might take a bit more time than 2 min!")
	}
}
