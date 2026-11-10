package main

import "fmt"

func main() {
	desserts := [2]string{"Cupcake", "Eclair"}
	slice := desserts[:1]

	fmt.Println("Length of slice:", len(slice), "Capacity of slice:", cap(slice), "Slice:", slice)
	slice = append(slice, "Macaron")
	fmt.Println(desserts)

	fmt.Println("Length of slice:", len(slice), "Capacity of slice:", cap(slice), "Slice:", slice)
	slice = append(slice, "Cake")
	fmt.Println(desserts)

	fmt.Println("Length of slice:", len(slice), "Capacity of slice:", cap(slice), "Slice:", slice)
	slice = append(slice, "Tiramisu")
	fmt.Println(desserts)

	fmt.Println("Length of slice:", len(slice), "Capacity of slice:", cap(slice), "Slice:", slice)
	slice = append(slice, "Chocolate cake")
	fmt.Println(desserts)

	fmt.Println("Length of slice:", len(slice), "Capacity of slice:", cap(slice), "Slice:", slice)

}
