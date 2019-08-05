package main

import (
	"bytes"
	"testing"
)

func TestPrintState(t *testing.T) {

	stateTests := []struct {
		name string
		grid Grid
		want string
	}{
		{
			name: "some taken",
			grid: Grid{
				{" ", "X", "X"},
				{"O", " ", "O"},
				{"X", "X", "O"},
			},
			want: "---------\n" +
				"|   X X |\n" +
				"| O   O |\n" +
				"| X X O |\n" +
				"---------\n",
		},
	}

	for _, tt := range stateTests {
		t.Run(tt.name, func(t *testing.T) {
			assertWriterOutput(t, tt.grid, tt.want)
		})
	}
}

func TestFromString(t *testing.T) {
	cells := "\"O OXXO XX\""
	grid := Grid{}
	grid.FromString(cells)
	want := "---------\n" +
		"| O   O |\n" +
		"| X X O |\n" +
		"|   X X |\n" +
		"---------\n"
	assertWriterOutput(t, grid, want)
}

func assertWriterOutput(t *testing.T, grid Grid, want string) {
	t.Helper()

	buffer := bytes.Buffer{}
	grid.PrintState(&buffer)
	got := buffer.String()

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
