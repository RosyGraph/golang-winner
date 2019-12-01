// Go supports anonymous functions, which can form
// closures. Anonymous functions are useful when you want
// to define a function inline without having to name it.
package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	for i := 0; i < 4; i++ {
		fmt.Printf("%d\n", nextInt())
	}

	newInts := intSeq()
	fmt.Println(newInts())
}