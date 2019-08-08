package main

import (
	"bytes"
	"testing"
)

func TestPrintState(t *testing.T) {

	stateTests := []struct {
		name       string
		gridString string
		want       string
	}{
		{name: "X wins", gridString: `"XXXO OXXO"`, want: "X wins"},
		{name: "O wins", gridString: `"OXXOOXXXO"`, want: "O wins"},
		{name: "draw", gridString: `"XXOOOXXXO"`, want: "Draw"},
		{name: "not finished", gridString: `"         "`, want: "Game not finished"},
		{name: "too many Xs", gridString: `"XXXOXOXXO"`, want: "Impossible"},
		{name: "both won", gridString: `" XOOXOXXO"`, want: "Impossible"},
	}

	for _, tt := range stateTests {
		t.Run(tt.name, func(t *testing.T) {
			grid := Grid{}
			grid.FromString(tt.gridString)
			got := grid.GameState()
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
		err := grid.HumanMove("1 1", "X")
		want := "---------\n" +
			"|       |\n" +
			"|       |\n" +
			"| X     |\n" +
			"---------\n"
		assertWriterOutput(t, grid, want)

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("cell is occupied", func(t *testing.T) {
		grid := Grid{}
		grid.FromString("\"X        \"")
		err := grid.HumanMove("1 3", X)

		if err != ErrCellOccupied {
			t.Errorf("wanted %q error but didn't get one", ErrCellOccupied)
		}
	})

	t.Run("input is not 2 numbers", func(t *testing.T) {
		grid := Grid{}
		grid.FromString("\"         \"")
		err := grid.HumanMove("a 1", X)

		if err != ErrInvalidInput {
			t.Errorf("wanted %q error but didn't get one", ErrInvalidInput)
		}
	})

	t.Run("input outside 1 to 3", func(t *testing.T) {
		grid := Grid{}
		grid.FromString("\"         \"")
		err := grid.HumanMove("4 5", X)

		if err != ErrInvalidInput {
			t.Errorf("wanted %q error but didn't get one", ErrInvalidInput)
		}
	})
}

func TestEasyMove(t *testing.T) {
	g := Grid{}
	g.FromString("\"  XO  OX \"")
	g.EasyMove(X)

	want := 3
	xCount := 0
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if g[row][col] == X {
				xCount++
			}
		}
	}

	if xCount != want {
		t.Errorf("wanted %d X's got %d", want, xCount)
	}
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
