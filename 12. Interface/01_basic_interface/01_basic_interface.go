package main

import "fmt"

type CoffeeMachine interface {
	Brew() string
	Clean() string
}

type CapsuleMachine struct {
	Brand string
}

func (c CapsuleMachine) Brew() string {
	return fmt.Sprintf("%s has brewed one of coffee", c.Brand)
}

func (c CapsuleMachine) Clean() string {
	return fmt.Sprintf("%s cleaned machine", c.Brand)
}

func main() {
	var machine CoffeeMachine

	machine = CapsuleMachine{
		Brand: "Nespresso",
	}
	fmt.Println(machine.Brew())
	fmt.Println(machine.Clean())
}
