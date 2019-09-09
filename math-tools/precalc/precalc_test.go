package main

import "testing"

func TestRoots(t *testing.T) {
	t.Run("two rational roots", func(t *testing.T) {
		got := Roots(1.0, -2.0, -3.0)
		want := [2]float64{3, -1}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("two rational roots", func(t *testing.T) {
		got := Roots(1.0, 2.0, -3.0)
		want := [2]float64{1, -3}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestParseCoefficients(t *testing.T) {
	got := ParseCoefficients("x^2 - 2x - 3")
	want := [3]float64{1, -2, -3}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
