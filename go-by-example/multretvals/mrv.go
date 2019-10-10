package main

import "fmt"

func vals() (int, int) {
	return 4, 5
}

func main() {
	fmt.Println(vals())

	_, c := vals()
	fmt.Printf("c: %d\n", c)
}
