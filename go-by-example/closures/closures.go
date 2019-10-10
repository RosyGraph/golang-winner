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
