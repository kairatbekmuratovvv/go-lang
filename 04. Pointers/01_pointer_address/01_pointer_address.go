package main

import "fmt"

func main() {
	coffee := "Americano"
	pointer := &coffee
	fmt.Println("Coffee name is: ", coffee)
	fmt.Println("Location on memory is: ", &coffee)
	fmt.Printf("Pointer address is: %p\n", pointer)

	fmt.Println("-----------------------------------------")

	coffeeCopy := coffee //value is copied to another location
	pointerCopy := &coffeeCopy
	fmt.Println("Coffee name is: ", coffeeCopy)
	fmt.Println("Location on memory is: ", &coffeeCopy)
	fmt.Printf("Pointer address is: %p\n", pointerCopy)
}
