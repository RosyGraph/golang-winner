package main

import "fmt"

func main() {
	var (
		x, y float64
		z    int
	)
	fmt.Println("This program illustrates a chaotic function.")
	fmt.Print("Enter a number between 0 and 1: ")
	fmt.Scan(&x)
	fmt.Print("Enter another number between 0 and 1: ")
	fmt.Scan(&y)
	fmt.Print("Enter the range: ")
	fmt.Scan(&z)
	fmt.Printf("%-5v %15v %10v\n", "Index", x, y)
	fmt.Println("------------------------------")
	for i := 0; i < z; i++ {
		x = 3.9 * x * (1 - x)
		y = 3.9 * y * (1 - y)
		fmt.Printf("%-10v%10.5f%10.5f\n", i+1, x, y)
	}
}
