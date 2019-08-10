package main

import "fmt"

func IntToRoman(num int) (ans string) {
	if num == 4 {
		ans = "IV"
	}

	if num == 9 {
		ans = "IX"
	}

	if num > 9 && num < 14 {
		ans = "X"
		num -= 10
	}

	if num > 4 && num < 9 {
		ans = "V"
		num -= 5
	}

	for ; num > 0 && num < 4; num-- {
		ans += "I"
	}
	return
}

func main() {
	fmt.Println("vim-go")
}
