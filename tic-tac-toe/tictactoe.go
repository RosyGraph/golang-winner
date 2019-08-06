package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const X = "X"
const O = "O"

type State struct {
	total          int
	xCount         int
	oCount         int
	xLeftDiagonal  int
	oLeftDiagonal  int
	xRightDiagonal int
	oRightDiagonal int
	xWins          bool
	oWins          bool
}

type Grid [3][3]string

func (g Grid) Print(writer io.Writer) {
	fmt.Fprintln(writer, "---------")
	for _, row := range g {
		fmt.Fprintf(writer, "| %s %s %s |\n", row[0], row[1], row[2])
	}
	fmt.Fprintln(writer, "---------")
}

func (grid Grid) GameState() string {
	s := State{}
	for row := 0; row < 3; row++ {
		xInRow := 0
		oInRow := 0
		xInCol := 0
		oInCol := 0

		grid.checkDiagonals(row, &s)

		for col := 0; col < 3; col++ {
			switch grid[row][col] {
			case X:
				xInRow++
				if xInRow > 2 {
					if s.oWins {
						return "Impossible"
					} else {
						s.xWins = true
					}
				}
			case O:
				oInRow++
				if oInRow > 2 {
					if s.xWins {
						return "Impossible"
					} else {
						s.oWins = true
					}
				}
			}

			switch grid[col][row] {
			case X:
				xInCol++
				s.xCount++
				if xInCol > 2 {
					if s.oWins {
						return "Impossible"
					} else {
						s.xWins = true
					}
				}
			case O:
				oInCol++
				s.oCount++
				if oInCol > 2 {
					if s.xWins {
						return "Impossible"
					} else {
						s.oWins = true
					}
				}
			}
		}
	}

	if abs(s.xCount-s.oCount) > 2 {
		return "Impossible"
	}
	s.total += s.xCount + s.oCount
	if s.xWins {
		return "X wins"
	} else if s.oWins {
		return "O wins"
	} else if s.total < 9 {
		return "Game not finished"
	} else {
		return "Draw"
	}
}

func (g *Grid) FromString(cells string) {
	cells = strings.Trim(cells, "\"")
	for i, char := range cells {
		g[i/3][i%3] = string(char)
	}
}

func (g Grid) checkDiagonal(diagonal *int, wins *bool) {
	*diagonal++
	if *diagonal > 2 {
		*wins = true
	}
}

func (grid Grid) checkDiagonals(row int, s *State) {
	switch grid[row][row] {
	case X:
		grid.checkDiagonal(&s.xLeftDiagonal, &s.xWins)
	case O:
		grid.checkDiagonal(&s.oLeftDiagonal, &s.oWins)
	}

	switch grid[row][2-row] {
	case X:
		grid.checkDiagonal(&s.xRightDiagonal, &s.xWins)
	case O:
		grid.checkDiagonal(&s.oRightDiagonal, &s.oWins)
	}
}

func main() {
	grid := Grid{}
	grid.Print(os.Stdout)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
