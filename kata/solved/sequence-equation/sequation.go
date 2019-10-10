package main

import "fmt"

func permutationEquation(p []int32) []int32 {
	m := map[int32]int32{}
	res := make([]int32, len(p))
	for i, v := range p {
		m[v] = int32(i) + 1
	}
	for i, _ := range p {
		j := m[int32(i+1)]
		res[i] = m[j]
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
