package music

import "testing"

func TestNoteFromString(t *testing.T) {
	got := NoteFromString("A" + Sharp)
	want := Note{"A", Sharp}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
