// Interfaces are named collections of method signatures.
package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return math.Pi * c.radius * 2
}

func measure(g geometry) {
	fmt.Printf("%v:\narea: %f\nperim: %f\n", g, g.area(), g.perim())
}

func main() {
	r := rect{width: 2, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
