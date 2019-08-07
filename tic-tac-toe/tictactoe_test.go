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
				{" ", "X", "X"},
				{"X", "O", "O"},
				{" ", "X", "O"},
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

func TestFromString(t *testing.T) {
	t.Run("mix", func(t *testing.T) {
		cells := "\"  OXXO  X\""
		grid := Grid{}
		grid.FromString(cells)
		want := "---------\n" +
			"|     O |\n" +
			"| X X O |\n" +
			"|     X |\n" +
			"---------\n"
		assertWriterOutput(t, grid, want)
	})

	t.Run("full", func(t *testing.T) {
		cells := "\"OOOXXOOXX\""
		grid := Grid{}
		grid.FromString(cells)
		want := "---------\n" +
			"| O O O |\n" +
			"| X X O |\n" +
			"| O X X |\n" +
			"---------\n"
		assertWriterOutput(t, grid, want)
	})

	t.Run("empty", func(t *testing.T) {
		cells := "\"         \""
		grid := Grid{}
		grid.FromString(cells)
		want := "---------\n" +
			"|       |\n" +
			"|       |\n" +
			"|       |\n" +
			"---------\n"
		assertWriterOutput(t, grid, want)
	})
}

func TestHumanMove(t *testing.T) {
	cells := "X X O    "
	grid := Grid{}
	grid.FromString(cells)
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
