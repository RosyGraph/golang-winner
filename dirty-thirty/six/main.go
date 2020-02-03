package main

import (
	"bufio"
	"fmt"
	"os"
)

func separate(s string) string {
	evens := ""
	odds := ""
	for index, c := range s {
		if index%2 == 0 {
			evens += string(c)
		} else {
			odds += string(c)
		}
	}
	return evens + " " + odds
}

func main() {
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	for buf.Scan() {
		line := buf.Text()
		fmt.Println(separate(line))
	}
}
