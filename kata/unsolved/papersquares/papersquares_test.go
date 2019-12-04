package main

import "testing"

func TestSolve(t *testing.T) {
	got := solve(3, 1)
	want := int64(2)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
