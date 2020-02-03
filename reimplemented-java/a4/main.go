package a4

import (
	"bufio"
	"strconv"
	"strings"
)

// Returns the smallest positive int in a string
func MinPositiveInt(s string) int {
	sc := wordScanner(s)
	sc.Scan()
	min, _ := strconv.Atoi(sc.Text())

	for sc.Scan() {
		current, _ := strconv.Atoi(sc.Text())
		if (current > 0) && (min == 0 || current < min) {
			min = current
		}
	}

	return min
}

// Returns the first word alphabetically
func FirstWordAlphabetically(s string) string {
	sc := wordScanner(s)
	sc.Scan()
	first := sc.Text()

	for sc.Scan() {
		current := sc.Text()
		if strings.Compare(current, first) < 0 {
			first = current
		}
	}

	return first
}

// Returns the smallest int in the two strings
func MinIntTwoStr(s1, s2 string) int {
	s := s1 + " " + s2
	sc := wordScanner(s)
	sc.Scan()
	min, _ := strconv.Atoi(sc.Text())

	for sc.Scan() {
		current, _ := strconv.Atoi(sc.Text())
		if current < min {
			min = current
		}
	}

	return min
}

// Curve the numbers in a space-separated String so the highest number is 100 and
// the difference between the highest number and 100 is added to each other number
func CurveScores(s string) string {
	sc := wordScanner(s)
	sc.Scan()
	max, _ := strconv.Atoi(sc.Text())
	var curved string

	for sc.Scan() {
		current, _ := strconv.Atoi(sc.Text())
		if current > max {
			max = current
		}
	}

	k := 100 - max
	sc = wordScanner(s)

	for sc.Scan() {
		current, _ := strconv.Atoi(sc.Text())
		current += k

		curved += strconv.FormatInt(int64(current), 10) + " "
	}

	return curved[:len(curved)-1]
}

// Returns a scanner with Split set to ScanWords
func wordScanner(s string) *bufio.Scanner {
	sc := bufio.NewScanner(strings.NewReader(s))
	sc.Split(bufio.ScanWords)

	return sc
}
