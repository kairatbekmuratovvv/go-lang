package main

import "fmt"

func calculateLoyaltyPoint(amount float64) int {
	var loyaltyPoint = int(amount * 2)
	return loyaltyPoint
}

func totalPointsUpdates(currentpoitn int, newEarnedPoint1 int) int {
	var totalPoint = currentpoitn + newEarnedPoint1
	return totalPoint

}

func main() {
	totalPoints := 120
	var newEarnedPoint int = calculateLoyaltyPoint(5.6)
	totalPointsUpdates(totalPoints, newEarnedPoint)
	fmt.Println(newEarnedPoint)
	fmt.Println(totalPointsUpdates(totalPoints, newEarnedPoint))

}
