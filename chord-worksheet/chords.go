package main

import "fmt"

const Sharp = '♯'
const Flat = '♭'
const Dblsharp = '𝄪'
const Dblflat = '𝄫'

func ValueOf(s string) (v int) {
	r := []rune(s)
	switch r[0] {
	case 'A':
		v = 0
	case 'B':
		v = 2
	case 'C':
		v = 3
	case 'D':
		v = 5
	case 'E':
		v = 7
	case 'F':
		v = 8
	case 'G':
		v = 10
	}
	if len(r) == 2 {
		switch r[1] {
		case Sharp:
			v++
		case Flat:
			v--
		case Dblsharp:
			v += 2
		case Dblflat:
			v += 2
		}
	}
	v %= 12
	return
}

func main() {
	fmt.Printf("%d %d", 'A', 'G')
}
