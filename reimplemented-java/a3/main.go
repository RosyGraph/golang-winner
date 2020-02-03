package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Returns a copy of the input string with each 'a' character censored.
func CensorA(s string) string {
	res := ""
	for _, c := range s {
		if c == 'a' || c == 'A' {
			c = '*'
		}
		res += string(c)
	}
	return res
}

// Returns true if there are more even numbers in the string than odd.
func ContainsMoreEvens(s string) bool {
	// NOTE: 0 is an even number
	sc := wordScanner(s)
	var evens int

	for sc.Scan() {
		token := sc.Text()
		n, err := strconv.Atoi(token)

		if err != nil {
			continue
		}

		if n%2 == 0 {
			evens++
		} else {
			evens--
		}
	}

	return evens > 0
}

// Returns a scanner with Split set to ScanWords
func wordScanner(s string) *bufio.Scanner {
	sc := bufio.NewScanner(strings.NewReader(s))
	sc.Split(bufio.ScanWords)

	return sc
}

func main() {
	fmt.Println("vim-go")
}
