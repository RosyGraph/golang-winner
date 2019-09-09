package main

import "fmt"

const Sharp = '♯'
const Flat = '♭'
const Dblsharp = '𝄪'
const Dblflat = '𝄫'

var alphabet = map[rune]int{
	'A': 0,
	'B': 2,
	'C': 3,
	'D': 5,
	'E': 7,
	'F': 8,
	'G': 10,
}

type Triad struct {
	root  string
	third string
	fifth string
}

func NoteValueOf(s string) (v int) {
	r := []rune(s)
	v = alphabet[r[0]]
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

func AscendingInt(r, i string) int {
	ival := NoteValueOf(i)
	rval := NoteValueOf(r)

	if ival < rval {
		ival += 12
	}

	return ival - rval
}

func main() {
	fmt.Printf("%d %d", 'A', 'G')
}
