package main

import "fmt"

func main() {
	var coffeeSizes = [3]string{"Small", "Medium", "Large"}
	fmt.Println(coffeeSizes)
	fmt.Println(len(coffeeSizes))
	coffeeSizes[len(coffeeSizes)-1] = "Extra Large"
	fmt.Println(coffeeSizes)
}
