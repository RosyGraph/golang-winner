package main

import "fmt"

type exponential func(x int) int

func main() {
	square := func(x int) int {
		return x * x
	}

	cube := func(x int) int {
		return x * x * x
	}

	flat := func(x int) int {
		return x
	}

	thrice := func(g exponential) int {
		return g(3)
	}

	twice := func(g exponential) int {
		return g(2)
	}

	which := func(x int) exponential {
		if 2*2 == x {
			return square
		} else if 2*2*2 == x {
			return cube
		}
		return flat
	}

	fmt.Println(thrice(square))
	fmt.Println(twice(square))
	fmt.Println(thrice(cube))
	fmt.Println(twice(cube))
	fmt.Println(which(8)(100))
}
