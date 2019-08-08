package main

import "fmt"

func main() {
	Greet("Elodie")
}

func Greet(s string) {
	fmt.Printf("Hello, %s", s)
}
