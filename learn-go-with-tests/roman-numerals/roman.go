package main

import "fmt"

const I = "I"
const V = "V"
const X = "X"

func IntToRoman(num int) (ans string) {
	if num == 19 {
		ans += "XIX"
		num -= 19
	}

	if num == 9 {
		ans = "IX"
		num -= 9
	}

	for num > 9 && num < 40 {
		ans += X
		num -= 10
	}

	if num == 4 {
		ans += "IV"
		num -= 4
	}

	if num > 4 && num < 9 {
		ans += V
		num -= 5
	}

	for ; num > 0 && num < 4; num-- {
		ans += I
	}
	return
}

func main() {
	fmt.Println("vim-go")
}
