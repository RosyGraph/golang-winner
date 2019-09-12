package main

import "testing"

func TestNoteFromString(t *testing.T) {
	tc := []struct {
		name string
		s    string
		want Note
	}{
		{name: "A sharp", s: "A" + Sharp, want: Note{"A", Sharp}},
		{name: "A natural", s: "A", want: Note{"A", Natural}},
	}
	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := NoteFromString(c.s)

			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}

func TestNoteValue(t *testing.T) {
	t.Run("value of A", func(t *testing.T) {
		got := NoteValue(Note{"A", Natural})
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("value of A sharp", func(t *testing.T) {
		got := NoteValue(Note{"A", Sharp})
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
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
		{name: "augmented 4th unison",
			root:     Note{"E", Natural},
			interval: Note{"A", Natural},
			want:     Interval{"perfect", 4}},
		{name: "diminished 7th",
			root:     Note{"C", Natural},
			interval: Note{"B", Dblflat},
			want:     Interval{"diminished", 7}},
		{name: "minor 7th",
			root:     Note{"D", Natural},
			interval: Note{"C", Natural},
			want:     Interval{"minor", 7}},
		{name: "major 3rd",
			root:     Note{"A", Natural},
			interval: Note{"C", Sharp},
			want:     Interval{"major", 3}},
		{name: "minor 3rd",
			root:     Note{"A", Natural},
			interval: Note{"C", Natural},
			want:     Interval{"minor", 3}},
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
