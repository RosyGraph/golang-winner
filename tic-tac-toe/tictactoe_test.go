package main

import (
	"bytes"
	"testing"
)

func TestPrintState(t *testing.T) {
	buffer := bytes.Buffer{}

	grid := Grid{
		{"X", "O", "X"},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	grid.PrintState(&buffer)

	got := buffer.String()
	want := "X O X \n      \n      \n"

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
