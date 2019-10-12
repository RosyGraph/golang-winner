// Go supports pointers, allowing you to pass references to
// values and records within your program.
package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Printf("initial: %v\n", i)

	zeroval(i)
	fmt.Printf("zeroval: %v\n", i)

	zeroptr(&i)
	fmt.Printf("zeroptr: %v\n", i)

	fmt.Printf("pointer: %v\n", &i)
}
