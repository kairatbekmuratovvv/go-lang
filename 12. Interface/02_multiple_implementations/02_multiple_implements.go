package main

import "fmt"

type CoffeeMachine interface {
	Brew() string
}

type DripMachine struct {
	//Brand string
	Model string
	Price int
}
type CapsuleMachine struct {
	Brand string
	Model string
	Price int
}

func (c CapsuleMachine) Brew() string {
	return fmt.Sprintf("%s  %s has brewed one of coffee", c.Brand, c.Model)
}

func (c DripMachine) Brew() string {
	return fmt.Sprintf("Drip coffee shot was prepared by %s", c.Model)
}

func (c DripMachine) Clean() string {
	return fmt.Sprintf("%s cleaned machine", c.Model)
}

func main() {
	var machineOne CoffeeMachine
	var machineTwo CoffeeMachine

	machineOne = CapsuleMachine{
		Brand: "Nespresso",
		Model: "XB23",
		Price: 199,
	}
	machineTwo = DripMachine{
		Model: "BrewPro",
		Price: 230,
	}
	var machineThree DripMachine

	fmt.Println(machineOne.Brew())
	fmt.Println(machineTwo.Brew())
	fmt.Println(machineThree.Clean())

}
