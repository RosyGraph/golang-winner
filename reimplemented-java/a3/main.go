package main

import (
	"fmt"
)

// Return a copy of the input string with
// each 'a' character replaced by '*'.
func CensorA(s string) string {
	res := ""
	for _, c := range s {
		if c == 'a' || c == 'A' {
			c = '*'
		}
		res += string(c)
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
