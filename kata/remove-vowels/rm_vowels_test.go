package main

import "testing"

func TestRemoveVowels(t *testing.T) {
	t.Run("basic case", func(t *testing.T) {
		s := "hello world"
		got := RemoveVowels(s)
		want := "hll wrld"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
