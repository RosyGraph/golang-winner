// for is Go's only looping construct. Here are three basic
// types of for loops.
package main

import "fmt"

func main() {
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("looping...")
		if i > 6 {
			break
		}
		i++
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
