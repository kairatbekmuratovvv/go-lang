package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)
	slice1 = append(slice1, 5)
	fmt.Println(slice1)
}
