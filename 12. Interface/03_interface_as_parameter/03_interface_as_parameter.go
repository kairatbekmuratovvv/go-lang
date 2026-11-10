package main

import "fmt"

type Barista interface {
	PrepareCoffee() string
}

type SeniorBarista struct {
	Name string
}

func (s SeniorBarista) PrepareCoffee() string {
	return fmt.Sprintf("%s prepared a caramel latte", s.Name)
}

type JuniorBarista struct {
	Name string
}

func (j JuniorBarista) PrepareCoffee() string {
	return fmt.Sprintf("%s made a hot chocolate", j.Name)
}
func ServeDrink(b Barista) {
	fmt.Println(b.PrepareCoffee())
	fmt.Println("Barista served coffee to client")
}

func main() {
	bogdan := SeniorBarista{Name: "Bogdan"}
	kiro := JuniorBarista{Name: "Kiro"}
	//fmt.Println(bogdan.PrepareCoffee())
	//fmt.Println(kiro.PrepareCoffee())
	ServeDrink(bogdan)
	ServeDrink(kiro)

}
