package main

import "math"

// Find the minimum integer height of a triangle given area and len of base
func lowestTriangle(b, a int) int {
	h := math.Ceil(float64(a) * 2.0 / float64(b))
	return int(h)
}
