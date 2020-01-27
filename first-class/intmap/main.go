package main

import "fmt"

func intMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func main() {
	a := []int{1, 2, 3, 4}
	r := intMap(a, func(n int) int {
		return n * 5
	})
	fmt.Println(r)
}
