package main

import "fmt"

func time(qutyCoffee int, secondPer int) (totalTime int) {
	totalTime = qutyCoffee * secondPer
	return
}

func main() {
	fmt.Println(time(20, 12))

}
