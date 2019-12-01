package main

import (
	"fmt"

	"github.com/RosyGraph/kata/solved/fibonacci/fibonacci"
)

func main() {
	var fib fibonacci.Fibonacci
	fmt.Printf("%d\n", fib.Fib(8))
}
