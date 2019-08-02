// This program accepts coin amounts and prints total cash value to console.
package main

import (
	"fmt"
	"strings"
)

func getTotal(c map[string]float64) float64 {
	var total float64
	total += c["pennies"]
	total += c["nickels"] * 5
	total += c["dimes"] * 10
	total += c["quarters"] * 25
	return total / 100
}

func getAmounts() map[string]float64 {
	c := make(map[string]float64)
	var v float64
	coinTypes := []string{"pennies", "nickels", "dimes", "quarters"}
	fmt.Println("Enter your coin amounts.")
	for _, i := range coinTypes {
		fmt.Printf("%v: ", strings.Title(i))
		fmt.Scan(&v)
		c[i] = v
	}
	return c
}

func main() {
	c := getTotal(getAmounts())
	fmt.Printf("Total: $%.02f\n", c)
}
