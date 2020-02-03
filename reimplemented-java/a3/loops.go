package a3

import (
	"bufio"
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
	var evens int

	for sc := wordScanner(s); sc.Scan(); {
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

// Returns a string resembling a triangle with height and width = n
func TriangleString(n int) string {
	res := ""

	for row := 0; row < n; row++ {
		for col := 0; col < row+1; col++ {
			res += "*"
		}
		res += "\n"
	}

	return res[:len(res)-1]
}

// Returns a color safe value between 0 and 255
func SafeColor(n int) int {
	switch {
	case n < 0:
		return 0
	case n > 255:
		return 255
	}
	return n
}

// Returns a scanner with Split set to ScanWords
func wordScanner(s string) *bufio.Scanner {
	sc := bufio.NewScanner(strings.NewReader(s))
	sc.Split(bufio.ScanWords)

	return sc
}
