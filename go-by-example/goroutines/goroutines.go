// A goroutine is a lightweight thread of execution.
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s: %d\n", from, i)
	}
}

func main() {
	f("direct") // Run synchronously

	go f("goroutine")

	go func(msg string) { // Goroutine for anonymous func
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}