package main

import "fmt"

func add(x, y float32) float32 {
	return x + y
}

func main() {
	num1, num2 := 5.6, 9.5
	fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "there"

	fmt.Println(w1, w2)
}
