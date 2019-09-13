package main

import "fmt"

const Sharp = "♯"
const Flat = "♭"
const Dblsharp = "𝄪"
const Dblflat = "𝄫"
const Natural = "♮"

var NoteValues = map[string]int{
	"A": 0,
	"B": 2,
	"C": 3,
	"D": 5,
	"E": 7,
	"F": 8,
	"G": 10,
}

var Alphabet = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
}

var AccidentalValues = map[string]int{
	"♯": 1,
	"♭": -1,
	"𝄪": 2,
	"𝄫": -2,
	"♮": 0,
}

var imperfect = []string{"diminished", "minor", "major", "augmented"}
var perfect = []string{"diminished", "perfect", "augmented"}

type Note struct {
	name       string
	accidental string
}

func NoteFromString(s string) Note {
	r := []rune(s)
	var accidental string
	name := string(r[0])

	if len(r) > 1 {
		accidental = string(r[1])
	} else {
		accidental = Natural
	}
	return Note{name, accidental}
}

func NoteValue(n Note) int {
	return (NoteValues[n.name] + AccidentalValues[n.accidental]) % 12
}

type Interval struct {
	quality  string
	quantity int
}

func AscendingInterval(r, i Note) Interval {
	var quantity int
	var quality string
	rval := Alphabet[r.name]
	ival := Alphabet[i.name]
	if rval > ival {
		ival += 7
	}
	quantity = ival - rval + 1
	fmt.Println(quantity)
	rcval := NoteValue(r)
	icval := NoteValue(i)
	if icval < rcval {
		icval += 12
	}
	totalv := icval - rcval
	// TODO: Refactor these ugly switch cases
	switch quantity {
	case 1:
		switch totalv {
		case 11:
			quality = "diminished"
		case 0:
			quality = "perfect"
		case 1:
			quality = "augmented"
		default:
			quality = "undefined"
		}
	case 2:
		quality = imperfect[totalv]
	case 3:
		quality = imperfect[totalv-2]
	case 4:
		quality = perfect[totalv-4]
	case 5:
		quality = perfect[totalv-6]
	case 6:
		quality = imperfect[totalv-7]
	case 7:
		quality = imperfect[totalv-9]
	}
	return Interval{quality, quantity}
}

func main() {
	d := Note{"D", Natural}
	c := Note{"C", Natural}
	fmt.Println(NoteValue(c) - NoteValue(d))
}
