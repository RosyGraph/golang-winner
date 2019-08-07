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
			name: "X wins",
			grid: Grid{
				{"X", "X", "X"},
				{"O", " ", "O"},
				{"X", "X", "O"},
			},
			want: "X wins",
		},
		{
			name: "O wins",
			grid: Grid{
				{"O", "X", "X"},
				{"O", "O", "X"},
				{"X", "X", "O"},
			},
			want: "O wins",
		},
		{
			name: "draw",
			grid: Grid{
				{"X", "X", "O"},
				{"O", "O", "X"},
				{"X", "X", "O"},
			},
			want: "Draw",
		},
		{
			name: "not finished",
			grid: Grid{
				{" ", " ", " "},
				{" ", " ", " "},
				{" ", " ", " "},
			},
			want: "Game not finished",
		},
		{
			name: "too many Xs",
			grid: Grid{
				{"X", "X", "X"},
				{"O", "X", "O"},
				{"X", "X", "O"},
			},
			want: "Impossible",
		},
		{
			name: "both won",
			grid: Grid{
				{" ", "X", "O"},
				{"O", "X", "O"},
				{"X", "X", "O"},
			},
			want: "Impossible",
		},
	}

	for _, tt := range stateTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.grid.GameState()
			if got != tt.want {
				t.Errorf("got %s want %s", got, tt.want)
			}
		})
	}
}

func TestHumanMove(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		grid := Grid{}
		grid.FromString("\"         \"")
		grid.HumanMove("1 1", "X")
		want := "---------\n" +
			"|       |\n" +
			"|       |\n" +
			"| X     |\n" +
			"---------\n"
		assertWriterOutput(t, grid, want)
	})

	t.Run("cell is occupied", func(t *testing.T) {
		grid := Grid{}
		grid.FromString("\"X        \"")
		err := grid.HumanMove("1 3", "X")

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})

	t.Run("input is not 2 numbers", func(t *testing.T) {
		grid := Grid{}
		grid.FromString("\"         \"")
		grid.HumanMove("1 1", "X")
		want := "---------\n" +
			"|       |\n" +
			"|       |\n" +
			"| X     |\n" +
			"---------\n"
		assertWriterOutput(t, grid, want)
	})

	t.Run("input outside 1 to 3", func(t *testing.T) {
		grid := Grid{}
		grid.FromString("\"         \"")
		grid.HumanMove("1 1", "X")
		want := "---------\n" +
			"|       |\n" +
			"|       |\n" +
			"| X     |\n" +
			"---------\n"
		assertWriterOutput(t, grid, want)
	})
}

func assertWriterOutput(t *testing.T, grid Grid, want string) {
	t.Helper()

	buffer := bytes.Buffer{}
	grid.Print(&buffer)
	got := buffer.String()

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
