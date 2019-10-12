// Functions are central in Go. We'll learn about functions
// with a few different examples.
package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Printf("1 + 2 = %d\n", res)

	res = plusPlus(1, 2, 3)
	fmt.Printf("1 + 2 + 3 = %d\n", res)

	plusPlusPlus := func(a, b, c, d int) int {
		return a + b + c + d
	}

	fmt.Printf("1 + 2 + 3 + 4 = %d\n",
		plusPlusPlus(1, 2, 3, 4))
}
