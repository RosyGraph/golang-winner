package main

import (
	"fmt"
	"os"
	"strconv"
)

func collatz(n int) int {
	var counter int

	for n > 1 {
		if n%2 != 0 {
			n = 3*n + 1
			counter++
		}
		counter++
		n /= 2
	}
	return counter
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < n+1; i++ {
		fmt.Printf("%d: %d steps\n", i, collatz(i))
	}
}
