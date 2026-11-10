package main

import "fmt"

func main() {
	desertMenu := [4]string{"Muffin", "Cheesecake", "Brownie", "Cookie"}
	fmt.Println(desertMenu)
	slice := desertMenu[1:3]
	fmt.Println(slice)
	slice = desertMenu[:]
	fmt.Println(slice)
}
