package main

import "fmt"

func weird(n int32) string {
	if n%2 != 0 {
		return "Weird"
	}
	if n > 5 && n < 21 {
		return "Weird"
	}
	return "Not Weird"
}

func main() {
	fmt.Println("vim-go")
}
