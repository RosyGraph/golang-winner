package music

const Sharp = "♯"
const Flat = "♭"
const Dblsharp = "𝄪"
const Dblflat = "𝄫"

type Note struct {
	name       string
	accidental string
}

func NoteFromString(s string) Note {
	return Note{"A", Sharp}
}
