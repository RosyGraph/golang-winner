// Channels are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another goroutine.
package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	fmt.Println("vim-go")
	msg := <-messages
	fmt.Println(msg)
}
