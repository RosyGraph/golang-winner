package iteration

import "testing"

func TestHelloWorld(t *testing.T) {
	got := repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got: '%s' want '%s'", got, want)
	}
}
