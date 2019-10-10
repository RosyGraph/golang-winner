// Branching with if and else in Go is straight-forward.
package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("Seven is even.")
	} else {
		fmt.Println("Seven is odd.")
	}

	if 8%4 == 0 {
		fmt.Println("Eight is divisible by four.")
	}

	if n := 9; n < 0 {
		fmt.Printf("%d is negative.\n", n)
	} else if n < 10 {
		fmt.Printf("%d has one digit.\n", n)
	} else {
		fmt.Printf("%d has multiple digits.\n", n)
	}
}
