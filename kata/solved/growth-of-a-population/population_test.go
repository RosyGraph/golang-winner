package kata

import "testing"

func TestNbYear(t *testing.T) {
	got := NbYear(1500, 5, 100, 5000)
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
