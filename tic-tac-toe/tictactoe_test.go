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
			buffer := bytes.Buffer{}
			tt.grid.PrintState(&buffer)
			got := buffer.String()
			assertString(t, got, tt.want)
		})
	}
}

func TestFromString(t *testing.T) {
	cells := "\"O OXXO XX\""
	grid := Grid{}
	grid.FromString(cells)
	buffer := bytes.Buffer{}
	grid.PrintState(&buffer)
	got := buffer.String()
	want := "---------\n" +
		"| O   O |\n" +
		"| X X O |\n" +
		"|   X X |\n" +
		"---------\n"
	assertString(t, got, want)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
