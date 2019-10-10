// Go supports methods defined on struct types.
package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 5}

	fmt.Printf("area:  %v\n", r.area())
	fmt.Printf("perim: %v\n", r.perim())

	rp := &r
	fmt.Printf("area:  %v\n", rp.area())
	fmt.Printf("perim: %v\n", rp.perim())
}
