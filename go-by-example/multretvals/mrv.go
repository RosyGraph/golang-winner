// Go has built-in support for multiple return values. This
// feature is used often in idiomatic Go, for example to return
// both result and error values from a function.
package main

import "fmt"

func vals() (int, int) {
	return 4, 5
}

func main() {
	fmt.Println(vals())

	_, c := vals()
	fmt.Printf("c: %d\n", c)
}
