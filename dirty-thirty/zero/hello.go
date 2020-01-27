package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)

	text, _ := buf.ReadString('\n')

	fmt.Println("Hello, World.\n" + text)
}
