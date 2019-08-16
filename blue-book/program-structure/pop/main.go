package main

import (
	"fmt"
	"time"

	"github.com/RosyGraph/blue-book/program-structure/popcount"
)

func main() {
	x := uint64(6)
	start := time.Now()
	pop := popcount.PopCount(x)
	fmt.Printf("PopCount(%v) = %v\n", x, pop)
	duration := time.Since(start)
	fmt.Printf("Took %.6f seconds.\n", duration.Seconds())
	start = time.Now()
	pop = popcount.PopCountIterative(x)
	fmt.Printf("PopCountIterative(%v) = %v\n", x, pop)
	duration = time.Since(start)
	fmt.Printf("Took %.6f seconds.\n", duration.Seconds())
}
