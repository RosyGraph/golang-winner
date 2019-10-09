package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func gcd(x, y int) int {
	for y != 0 {
		for y != 0 {
			x, y = y, x%y
		}
	}
	return x
}

func argErr() {
	log.Print("Unknown command. Usage: ./gcd [ int ] [ int ]\n")
	os.Exit(1)
}

func main() {
	n := os.Args[1:]
	if len(n) != 2 {
		argErr()
	}
	x, err := strconv.Atoi(n[0])
	if err != nil {
		argErr()
	}
	y, err := strconv.Atoi(n[1])
	if err != nil {
		argErr()
	}
	d := gcd(x, y)
	fmt.Println(d)
}
