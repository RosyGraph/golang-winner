package main

// TODO: Change intervals to reflect quantity and quality

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

type Interval struct {
	quality  string
	quantity int
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

func AscendingInt(r, i string) Interval {
	return Interval{}
}

func Quality(t Triad) string {
	third := AscendingInt(t.root, t.third)
	fifth := AscendingInt(t.root, t.fifth)

	switch {
	case third == 4 && fifth == 7:
		return "Major"
	case third == 3 && fifth == 7:
		return "Minor"
	case third == 3 && fifth == 6:
		return "Diminished"
	case third == 4 && fifth == 8:
		return "Augmented"
	}
	return "Not a triad"
}
