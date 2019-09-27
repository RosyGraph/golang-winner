package music

import (
	"bytes"
	"reflect"
	"testing"
)

func TestTriadEquals(t *testing.T) {
	t.Run("No sharps or flats", func(t *testing.T) {
		triad := TriadFromString("A C E")
		got := triad.Equals(TriadFromString("A C E"))
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestIntervalTest(t *testing.T) {
	buffer := bytes.Buffer{}
	IntervalTest(&buffer)
	got := buffer.String()
	want := "Identify the interval C to E\n"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestTriadFromString(t *testing.T) {
	got := TriadFromString("A C E")
	want := Triad{NoteFromString("A"), NoteFromString("C"), NoteFromString("E")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestTriadQuality(t *testing.T) {
	tc := []struct {
		name  string
		triad Triad
		want  string
	}{
		{
			name:  "C major",
			triad: Triad{NoteFromString("C"), NoteFromString("E"), NoteFromString("G")},
			want:  "major",
		},
		{
			name:  "A minor",
			triad: Triad{NoteFromString("A"), NoteFromString("C"), NoteFromString("E")},
			want:  "minor",
		},
		{
			name:  "B diminished",
			triad: Triad{NoteFromString("B"), NoteFromString("D"), NoteFromString("F")},
			want:  "diminished",
		},
		{
			name: "G sharp augmented",
			triad: Triad{NoteFromString("G" + Sharp), NoteFromString("B" + Sharp),
				NoteFromString("D" + Dblsharp)},
			want: "augmented",
		},
	}
	for _, c := range tc {
		got := c.triad.Quality()
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestNoteValue(t *testing.T) {
	tc := []struct {
		name string
		note Note
		want int
	}{
		{name: "A value", note: NoteFromString("A"), want: 0},
		{name: "A natural value", note: NoteFromString("A" + Natural), want: 0},
		{name: "A sharp value", note: NoteFromString("A" + Sharp), want: 1},
		{name: "D flat value", note: NoteFromString("D" + Flat), want: 4},
	}
	for _, c := range tc {
		got := c.note.Value()
		if got != c.want {
			t.Errorf("got %d want %d", got, c.want)
		}
	}
}

func TestAscendingInterval(t *testing.T) {
	tc := []struct {
		name     string
		root     Note
		interval Note
		want     Interval
	}{
		{name: "perfect unison",
			root:     Note{"A", Natural},
			interval: Note{"A", Natural},
			want:     Interval{"perfect", 1}},
		{name: "diminished unison",
			root:     Note{"A", Natural},
			interval: Note{"A", Flat},
			want:     Interval{"diminished", 1}},
		{name: "augmented unison",
			root:     Note{"A", Natural},
			interval: Note{"A", Sharp},
			want:     Interval{"augmented", 1}},
		{name: "major 2nd",
			root:     Note{"E", Flat},
			interval: Note{"F", Natural},
			want:     Interval{"major", 2}},
		{name: "minor 3rd",
			root:     Note{"A", Natural},
			interval: Note{"C", Natural},
			want:     Interval{"minor", 3}},
		{name: "major 3rd",
			root:     Note{"A", Natural},
			interval: Note{"C", Sharp},
			want:     Interval{"major", 3}},
		{name: "augmented 4th",
			root:     Note{"E", Natural},
			interval: Note{"A", Sharp},
			want:     Interval{"augmented", 4}},
		{name: "diminished 5th",
			root:     Note{"C", Sharp},
			interval: Note{"G", Natural},
			want:     Interval{"diminished", 5}},
		{name: "major 6th",
			root:     Note{"F", Natural},
			interval: Note{"D", Natural},
			want:     Interval{"major", 6}},
		{name: "diminished 7th",
			root:     Note{"C", Natural},
			interval: Note{"B", Dblflat},
			want:     Interval{"diminished", 7}},
		{name: "minor 7th",
			root:     Note{"D", Natural},
			interval: Note{"C", Natural},
			want:     Interval{"minor", 7}},
		{name: "major 7th",
			root:     Note{"C", Natural},
			interval: Note{"B", Natural},
			want:     Interval{"major", 7}},
	}
	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := AscendingInterval(c.root, c.interval)

			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}
