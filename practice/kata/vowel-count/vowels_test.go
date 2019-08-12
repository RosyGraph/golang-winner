package kata

import "testing"

func TestGetCount(t *testing.T) {
	got := GetCount("hello world")
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
