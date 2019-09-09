package music

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

type Interval struct {
	quality  string
	quantity int
}

func AscendingInterval(r, i Note) Interval {
	var quantity int
	/* var quality string */
	rval := Alphabet[r.name]
	ival := Alphabet[i.name]
	if rval > ival {
		ival += 8
	}
	quantity = ival - rval + 1
	return Interval{"major", quantity}
}
