package main

import "fmt"

func Roots(p []int) [2]int {
	m := p[0] * p[2]
	var divisors []int
	if m < 0 {
		for i := -1; i >= m; i-- {
			if m%i == 0 && m/i+i == p[1] {
				divisors = append(divisors, i)
			}
		}
	} else {
		for i := 1; i <= m; i++ {
			if m%i == 0 && m/i+i == p[1] {
				divisors = append(divisors, i)
			}
		}
	}
	fmt.Println(divisors)
	return [2]int{divisors[0], divisors[1]}
}

func main() {
	fmt.Println("vim-go")
}
