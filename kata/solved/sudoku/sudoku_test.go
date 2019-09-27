package main

import (
	"testing"
)

func TestIsSolved(t *testing.T) {
	t.Run("9x9 solved", func(t *testing.T) {
		sudoku := `5 8 9 6 7 4 2 1 3
7 4 3 1 8 2 9 5 6
1 2 6 9 5 3 8 7 4
9 3 5 4 2 1 7 6 8
4 1 2 8 6 7 3 9 5
6 7 8 3 9 5 1 4 2
8 6 4 2 1 9 5 3 7
3 9 7 5 4 8 6 2 1
2 5 1 7 3 6 4 8 9
`
		got := IsSolved(sudoku)
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("2x2 unsolved", func(t *testing.T) {
		sudoku := `1 1 2 2
1 1 2 2
3 3 4 4
3 3 4 4
`
		got := IsSolved(sudoku)
		want := false
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
