package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Cleaning ...")
		fmt.Println("Suspending...")
	}()

	fmt.Println("Brewing fresh cup of espresso...")
	fmt.Println("Brewing fresh cup of latte...")

}
