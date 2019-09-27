package main

import "testing"

func TestRepeatedString(t *testing.T) {
	t.Run("standard case", func(t *testing.T) {
		got := repeatedString("aba", 10)
		want := int64(7)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("n = 1000000000000", func(t *testing.T) {
		got := repeatedString("a", 1000000000000)
		want := int64(1000000000000)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
