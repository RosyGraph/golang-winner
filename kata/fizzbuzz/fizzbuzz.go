// Kata from codingdojo.org
// Fizzbuzz prints the numbers from 1 - 100
// For multiples of three print "Fizz"
// For multiples of five print "Buzz"
// For multiples of three AND five print "FizzBuzz"
package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i < 101; i++ {
		fmt.Println(FizzBuzz(i))
	}
}

func FizzBuzz(number int) string {
	switch {
	case number%3 == 0 && number%5 == 0:
		return "FizzBuzz"
	case number%3 == 0:
		return "Fizz"
	case number%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(number)
	}
}
