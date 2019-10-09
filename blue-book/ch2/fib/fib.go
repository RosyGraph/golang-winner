package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	s := os.Args[1]
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Print("Usage: ./fib [ int ]")
	}
	fmt.Println(fib(n))
}
