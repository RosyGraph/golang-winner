package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/RosyGraph/mathlib"
)

func main() {
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	n, d := mathlib.Simplify(x, y)
	fmt.Printf("%d %d\n", n, d)
}
