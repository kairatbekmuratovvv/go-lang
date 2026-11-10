package main

import "fmt"

func greeting() (greetingMessage string) {
	var name string
	fmt.Scan(&name)
	greetingMessage = "Hello " + name
	return greetingMessage
}

func main() {
	fmt.Printf("%s\n", greeting())

}
