package main

import (
	"fmt"
)

func Greeting(coffeeShopName string, location string) {
	var greetingMessage = "Welcome to our coffee shop "
	fmt.Println(greetingMessage, coffeeShopName, location)
}

func main() {
	Greeting("Starbucks", "Abay")
	Greeting("Urbo", "")

}
