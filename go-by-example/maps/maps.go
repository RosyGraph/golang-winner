// Maps are Go's built-in associative data type (sometimes
// called hashes or dicts in other languages).
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["a"] = 7
	m["b"] = 13

	fmt.Printf("map: %v\n", m)
	fmt.Printf("v1:  %v\n", m["a"])

	fmt.Printf("len: %v\n", len(m))

	delete(m, "b")
	fmt.Printf("map: %v\n", m)

	_, ok := m["b"]
	fmt.Printf("ok:  %v\n", ok)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Printf("map: %v\n", n)
}
