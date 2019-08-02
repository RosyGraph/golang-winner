package main

import "fmt"

func addInterest(balances []float64, rate float64) []float64 {
	for i := 0; i < len(balances); i++ {
		balances[i] = balances[i] * (1 + rate)
	}
	return balances
}

func main() {
	principal := []float64{1000, 2200, 800, 360}
	rate := .05
	amounts := addInterest(principal, rate)
	for _, v := range amounts {
		fmt.Println(v)
	}
}
