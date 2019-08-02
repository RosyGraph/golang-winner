// This program shifts the given message left or right.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func getShift() int {
	var n int
	fmt.Print("Enter the shift (positive numbers shift forward): ")
	fmt.Scanf("%d", &n)
	return n
}

func getPhrase() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the phrase to be shifted: ")
	s, _ := reader.ReadString('\n')
	return s
}

func isLetter(n rune) bool {
	if n > 96 && n < 123 {
		return true
	}
	if n > 64 && n < 91 {
		return true
	}
	return false
}

// rune values {a:97 z:122 A:65 Z:90}
func cipher(p string, n int) {
	if n < 0 {
		n = 26 + (n % 26)
	}
	for _, c := range p {
		switch {
		case c > 96 && c < 123:
			fmt.Print(string(shiftLower(c, n)))
		case c > 64 && c < 91:
			fmt.Print(string(shiftUpper(c, n)))
		default:
			fmt.Print(string(c))
		}
	}
}

func shiftLower(c rune, s int) rune {
	return rune(((int(c) - 97 + s) % 26) + 97)
}

func shiftUpper(c rune, s int) rune {
	return rune(((int(c) - 65 + s) % 26) + 65)
}

func main() {
	p := getPhrase()
	s := getShift()
	cipher(p, s)
}
