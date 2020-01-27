package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var (
		userInt    uint64
		userFloat  float64
		userString string
	)

	// Read and save an integer, double, and String to your variables.
	scanner.Scan()
	userInt, _ = strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	userFloat, _ = strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	userString = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + userInt)

	// Print the sum of the double variables on a new line.
	sum := d + userFloat
	fmt.Printf("%.1f\n", sum)

	// Concatenate and print the String variables on a new line
	fmt.Println(s + userString)
	// The 's' variable above should be printed first.
}
