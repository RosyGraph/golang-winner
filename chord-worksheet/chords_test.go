package main

import "testing"

func TestValueOf(t *testing.T) {
	tc := []struct {
		name string
		note string
		want int
	}{
		{name: "note value of A", note: "A", want: 0},
		{name: "note value of C", note: "C", want: 3},
		{name: "note value of C♯", note: "C♯", want: 4},
		{name: "note value of D𝄪", note: "D𝄪", want: 7},
		{name: "note value of B𝄫", note: "B𝄫", want: 4},
		{name: "note value of E", note: "E", want: 7},
		{name: "note value of F♭", note: "F♭", want: 7},
		{name: "note value of F♭", note: "F♭", want: 7},
		{name: "note value of G𝄪", note: "G𝄪", want: 0},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := NoteValueOf(c.note)

			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}

func TestAscendingInt(t *testing.T) {
	got := AscendingInt("A", "C")
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

/* func TestQuality(t *testing.T) {
	triad := Triad{"C", "E", "G"}
	got := Quality(triad)
	want := "Major"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
} */
