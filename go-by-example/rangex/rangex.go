// Range iterates over elements in a variety of data
// structures. Let's see how using range with some of the
// data structures we've already learned.
package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, v := range nums {
		sum += v
	}

	fmt.Printf("sum: %7v\n", sum)

	for i, v := range nums {
		if v == 3 {
			fmt.Printf("3@index:%4d\n", i)
		}
	}

	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Printf("%s <-> %6s\n", k, v)
	}

	for k := range m {
		fmt.Printf("key: %7s\n", k)
	}

	for _, c := range "go" {
		fmt.Printf("%c: %9[1]d\n", c)
	}
}
