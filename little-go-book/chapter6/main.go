package main

import (
	"fmt"
	// "math/rand"
	// "sort"
)

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]
	fmt.Println(power, exists)
}

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	// swap the last value and the value we want to removes
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
