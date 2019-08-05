package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)
	got := buffer.String()

	// Print 3
	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	// Print 1, 2, 3 and Go!
	// Wait 1 second between each line
}
