package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s', want '%s'.", got, want)
		}
	}

	t.Run("Say hello to people.", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello world!' when an empty string is supplied.", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in Spanish.", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in French.", func(t *testing.T) {
		got := Hello("Antonin", "French")
		want := "Bonjour, Antonin!"
		assertCorrectMessage(t, got, want)
	})
}
