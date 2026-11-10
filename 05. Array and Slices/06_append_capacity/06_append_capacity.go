package main

import "fmt"

func main() {
	menu := []string{"Cake", "Pie"}

	fmt.Println("Initial menu:", menu)
	fmt.Println("length:", len(menu), "Capacity:", cap(menu))

	menu = append(menu, "Donut")
	fmt.Println("Menu after adding donut:", menu)
	fmt.Println("Menu after add", menu)
	fmt.Println("length:", len(menu), "Capacity:", cap(menu))
	fmt.Printf("Memory location: %p\n", &menu)

	menu = append(menu, "Ice cream")
	fmt.Println("length:", len(menu), "Capacity:", cap(menu))
	fmt.Printf("Memory location: %p\n", &menu)

	menu = append(menu, "Cheese cake")
	fmt.Println("length:", len(menu), "Capacity:", cap(menu))
	fmt.Printf("Memory location: %p\n", &menu)
	fmt.Printf("Memory location: %p\n", &menu[0])
	fmt.Printf("Memory location: %p\n", &menu[1])
	fmt.Printf("Memory location: %p\n", &menu[2])

}
