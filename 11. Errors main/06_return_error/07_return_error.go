package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("coffee_orders.txt")

	if err != nil {
		//fmt.Println("Error: could not open coffee orders file", err)
		//return
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist", err)
		} else {
			fmt.Println("General file opening error", err)
		}
		return
	}

	fmt.Println("Successfully open file:", file.Name())
}
