package main

import "testing"

// 1 + 1 + 1 + 1
func TestNthArithmetic(t *testing.T) {
	got := NthArithmetic(1, 1)
	want := 1

	if got != want {
		t.Errorf("NthArithmetic(%d, %d) got %d want %d", 1, 1, got, want)
	}
}
