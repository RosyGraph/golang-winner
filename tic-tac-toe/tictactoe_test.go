package main

import (
	"bytes"
	"testing"
)

func TestGrid(t *testing.T) {
	buffer := bytes.Buffer{}
	grid := Grid{
		{"X", "O", "X"},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	grid.PrintState(&buffer)

	want := "X O X\n      \n      "
	got := buffer.String()

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
