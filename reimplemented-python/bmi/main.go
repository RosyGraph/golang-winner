package main

import (
	"fmt"
	"math"
)

func getDimensions() (float64, float64) {
	// Get weight and height from std input
	var height, weight float64
	fmt.Printf("Enter your height in inches: ")
	fmt.Scan(&height)
	fmt.Printf("Enter your weight in lbs: ")
	fmt.Scan(&weight)
	return height, weight
}

func calculateBMI(height, weight float64) float64 {
	bmi := (weight * 720) / (height * height)
	return math.Round(bmi*100) / 100
}

func evalBMI(bmi float64) {
	fmt.Printf("The BMI is %v.\n", bmi)
	switch {
	case bmi < 19:
		fmt.Println("The BMI is below healthy range.")
	case bmi > 25:
		fmt.Println("The BMI is above healthy range.")
	default:
		fmt.Println("The BMI is within healthy range.")
	}
}

func main() {
	height, weight := getDimensions()
	bmi := calculateBMI(height, weight)
	evalBMI(bmi)
}
