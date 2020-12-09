package main

import (
	"fmt"
)

func Binary(n int) int {
	s := fmt.Sprintf("%b", n)
	c := 0
	max := 0
	flag := false
	for _, r := range s {
		if r != '1' {
			flag = false
			c = 0
			continue
		}
		flag = true
		if c++; flag {
			if c > max {
				max = c
			}
		}
	}
	return max
}

func main() {
	fmt.Println("vim-go")
}
