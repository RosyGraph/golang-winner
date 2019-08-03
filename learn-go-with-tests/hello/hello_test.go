package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got: %q Want: %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris."

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, world.' when an empty string is supplied.", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World."

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris."

		assertCorrectMessage(t, got, want)
	})
}
