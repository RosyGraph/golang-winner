package main

import "testing"

func TestValueOf(t *testing.T) {
	t.Run("value of A", func(t *testing.T) {
		got := ValueOf("A")
		want := 0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("value of C", func(t *testing.T) {
		got := ValueOf("C")
		want := 3

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("value of C♯", func(t *testing.T) {
		got := ValueOf("C♯")
		want := 3

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
