package main

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

type MockRandomizer struct{}

func (r *MockRandomizer) RandomChoice(arr [][2]int) [2]int {
	return arr[0]
}

func TestPlayGame(t *testing.T) {
	t.Run("human vs computer", func(t *testing.T) {
		var b bytes.Buffer
		b.WriteString("1 1\n2 2\n3 3\n")
		buffer := io.Reader(&b)
		randomizer := MockRandomizer{}
		got := PlayGame(buffer, &randomizer)
		want := "X wins"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestPrintState(t *testing.T) {
	stateTests := []struct {
		name       string
		gridString string
		want       string
	}{
		{name: "X wins row", gridString: `"XXXO OXXO"`, want: "X wins"},
		{name: "X wins col", gridString: `"XOOX OXX "`, want: "X wins"},
		{name: "O wins row", gridString: `"OOOXX X  "`, want: "O wins"},
		{name: "O wins col", gridString: `"OXXOOXXXO"`, want: "O wins"},
		{name: "draw", gridString: `"XXOOOXXXO"`, want: "Draw"},
		{name: "not finished", gridString: `"         "`, want: "Game not finished"},
		{name: "too many Xs", gridString: `"XXXOXOXXO"`, want: "Impossible"},
		{name: "both won", gridString: `" XOOXOXXO"`, want: "Impossible"},
		{name: "X win in row O win in row", gridString: `"XXXOOOX  "`, want: "Impossible"},
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
	g := Grid{{X, X, O}, {O, X, O}, {X, O, X}}

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
			var b bytes.Buffer
			b.WriteString(tt.move)
			buffer := io.Reader(&b)
			grid := Grid{}
			grid.FromString(tt.gridString)
			err := grid.HumanMove(buffer, tt.team)

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
	var randomizer DefaultRandomizer
	g.EasyMove(&randomizer, X)

	want := 3
	assertNum(t, g, X, want)
}

func assertNum(t *testing.T, g Grid, team string, want int) {
	t.Helper()

	got := 0
	for _, row := range g {
		for _, v := range row {
			if v == team {
				got++
			}
		}
	}

	if want != got {
		t.Fatalf("got %v %ves want %v", got, team, want)
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
