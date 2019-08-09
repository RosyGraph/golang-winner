package main

import (
	"bytes"
	"reflect"
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

func TestPrint(t *testing.T) {
	g := Grid{
		{X, X, O},
		{O, X, O},
		{X, O, X},
	}

	want := `---------
| X X O |
| O X O |
| X O X |
---------
`
	assertWriterOutput(t, g, want)
}

func TestHumanMove(t *testing.T) {
	humanTests := []struct {
		name       string
		gridString string
		move       string
		team       string
		err        error
		want       string
	}{
		{name: "valid input", gridString: `"         "`, move: "1 1",
			team: X, err: nil, want: `"      X  "`},
		{name: "cell is occupied", gridString: `"X        "`, move: "1 3",
			team: X, err: ErrCellOccupied, want: `"X        "`},
		{name: "input is not 2 numbers", gridString: `"         "`, move: "a 1",
			team: X, err: ErrInvalidInput, want: `"         "`},
		{name: "input outside 1 to 3", gridString: `"         "`, move: "4 5",
			team: O, err: ErrInvalidInput, want: `"         "`},
	}

	for _, tt := range humanTests {
		t.Run(tt.name, func(t *testing.T) {
			grid := Grid{}
			grid.FromString(tt.gridString)
			err := grid.HumanMove(tt.move, tt.team)

			if err != tt.err {
				t.Errorf("wanted %q got %q", tt.err, err)
			}

			want := Grid{}
			want.FromString(tt.want)
			if !reflect.DeepEqual(grid, want) {
				t.Errorf("wanted %v got %v", want, grid)
			}
		})
	}
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
