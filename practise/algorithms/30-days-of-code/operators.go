package main

import (
	"fmt"
	"math"
)

func main() {
	mealCost := 12.00
	tipPercent := 20
	taxPercent := 8

	tip := (mealCost * float64(tipPercent)) / 100
	tax := (mealCost * float64(taxPercent)) / 100
	totalCost := float64(mealCost) + tip + tax

	result := math.Round(totalCost)
	fmt.Println(result)
}
