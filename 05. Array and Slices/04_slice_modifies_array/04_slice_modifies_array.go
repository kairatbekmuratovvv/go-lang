package main

import "fmt"

func main() {
	menu := [3]string{"Tea", "Coffee", "Juice"}
	slice := menu[:2]
	fmt.Println("Before slice change: ")
	fmt.Println("menu: ", menu)
	fmt.Println("slice: ", slice)

	fmt.Println("length of menu array: ", len(menu))
	fmt.Println("length of slice: ", len(slice))

	slice[0] = "Match" // change 1 element slice
	fmt.Println("--------------------------------")
	fmt.Println("After slice change: ")
	fmt.Println("menu: ", menu) // changed
	fmt.Println("slice: ", slice)

}
