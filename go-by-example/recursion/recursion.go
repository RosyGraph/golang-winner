// Go supports recursive functions. Here's a classic factorial
// example.
package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Printf("%d\n", fact(4))
}
