package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Printf("Sum of two number is: %.", sum(a, b))
}
