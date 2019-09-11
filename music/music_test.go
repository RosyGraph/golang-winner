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
	t.Run("major 3rd", func(t *testing.T) {
		got := AscendingInterval(Note{"A", Natural}, Note{"C", Sharp})
		want := Interval{"major", 3}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("minor 3rd", func(t *testing.T) {
		got := AscendingInterval(Note{"A", Natural}, Note{"C", Natural})
		want := Interval{"minor", 3}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
