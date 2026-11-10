package main

import "fmt"

func info(customerName string, coffeeName string, price float32) {
	infoMessage := "%s 's favorite coffe is %s and price is  $%.3f\n"
	fmt.Printf(infoMessage, customerName, coffeeName, price)
}

func main() {
	info("Kiro", "Americano", 3.4)

}
