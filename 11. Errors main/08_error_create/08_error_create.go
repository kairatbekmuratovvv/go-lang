package main

import "fmt"

func main() {
	var err error
	err = fmt.Errorf("Some of error")

	if err == nil {
		fmt.Println("There is no error!")
	} else {
		fmt.Println("Error:", err)
	}
}
