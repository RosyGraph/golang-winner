package main

import "fmt"

func IntToRoman(num int) (ans string) {
	for ; num > 0; num-- {
		ans += "I"
	}
	return
}

func main() {
	fmt.Println("vim-go")
}
