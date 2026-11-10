package main

import "fmt"

func createTemp() (func(change float64) float64, float64) {
	baseTemp := 90.00
	adject := func(change float64) float64 {
		baseTemp = baseTemp + change
		return baseTemp
	}
	return adject, baseTemp
}

func main() {
	a, b := createTemp()
	fmt.Printf("Base temp %.1f\n", b)
	fmt.Printf("change is %.2f\n", a(-7.5))
	fmt.Printf("change is last %.2f\n", a(1))

}
