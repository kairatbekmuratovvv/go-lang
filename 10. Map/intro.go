package main

import (
	"fmt"
)

func main() {

	var country = map[int]string{
		1: "USA",
		2: "Turkey",
		4: "Korea",
	}

	fmt.Println(country[2])

	if val, ok := country[4]; ok {
		fmt.Println(val)
	}

	for key, value := range country {
		fmt.Println(value, " - ", key)

	}
	country[3] = "Russia"
	fmt.Println(country)

	for key, value := range country {
		fmt.Println("key: ", key, " value: ", value)
	}

	delete(country, 3)
	fmt.Println(country)

}
