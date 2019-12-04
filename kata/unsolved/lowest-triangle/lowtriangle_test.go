package main

import "testing"

func TestLowestTriangle(t *testing.T) {
	got := lowestTriangle(2, 2)
	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
