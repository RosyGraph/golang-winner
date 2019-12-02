package main

import (
	"fmt"
	"math"
)

func Roots(a, b, c float64) [2]float64 {
	root1 := (-b + math.Sqrt((b*b)-4*a*c)) / 2 * a
	root2 := (-b - math.Sqrt((b*b)-4*a*c)) / 2 * a
	return [2]float64{root1, root2}
}

func ParseCoefficients(s string) [3]float64 {
	return [3]float64{0, 0, 0}
}

func main() {
	fmt.Println("vim-go")
}
