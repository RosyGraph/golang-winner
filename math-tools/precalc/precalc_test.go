package main

import "testing"

func TestPolynomial(t *testing.T) {
	t.Run("2 rational roots", func(t *testing.T) {
		p := []int{1, 2, -3}
		got := Roots(p)
		want := [2]int{1, -3}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("2 rational roots", func(t *testing.T) {
		p := []int{1, -2, -3}
		got := Roots(p)
		want := [2]int{-1, 3}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
