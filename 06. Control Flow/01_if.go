package main

import "fmt"

func main() {
	orderTotal := 15.0

	if orderTotal > 10 {
		fmt.Println("You got free cookies!")
	}

	orderTotal = 10
	if orderTotal > 10 {
		fmt.Println("You can't !")
	}

}
