package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

type XOCounts struct {
	xInRow int
	oInRow int
	xInCol int
	oInCol int
}

type Grid [3][3]string

func (g *Grid) HumanMove(move, team string) {
	coords := strings.Split(move, " ")
	row, _ := strconv.Atoi(coords[1])
	col, _ := strconv.Atoi(coords[0])
	row = 3 - row
	col -= 1
	g[row][col] = team
}

func (g Grid) GameState() string {
	s := State{}
	for row := 0; row < 3; row++ {
		counts := XOCounts{}

		g.checkDiagonals(row, &s)

		for col := 0; col < 3; col++ {
			if !g.checkRows(&s, &counts, row, col) || !g.checkCols(&s, &counts, row, col) {
				return "Impossible"
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

func (g Grid) Print(writer io.Writer) {
	fmt.Fprintln(writer, "---------")
	for _, row := range g {
		fmt.Fprintf(writer, "| %s %s %s |\n", row[0], row[1], row[2])
	}
	fmt.Fprintln(writer, "---------")
}

func (g *Grid) FromString(cells string) {
	cells = strings.ReplaceAll(cells, "\"", "")
	for i, char := range cells {
		if char == ' ' {
			g[i/3][i%3] = " "
		} else {
			g[i/3][i%3] = string(char)
		}
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

func (g Grid) checkRow(s *State, inRow *int, oppWins bool, wins *bool) bool {
	*inRow++
	if *inRow > 2 {
		if oppWins {
			return false
		} else {
			*wins = true
		}
	}
	return true
}

func (g Grid) checkRows(s *State, counts *XOCounts, row, col int) bool {
	switch g[row][col] {
	case X:
		if !g.checkRow(s, &counts.xInRow, s.oWins, &s.xWins) {
			return false
		}
	case O:
		if !g.checkRow(s, &counts.oInRow, s.xWins, &s.oWins) {
			return false
		}
	}
	return true
}

func (g Grid) checkCols(s *State, counts *XOCounts, row, col int) bool {
	switch g[col][row] {
	case X:
		counts.xInCol++
		s.xCount++
		if counts.xInCol > 2 {
			if s.oWins {
				return false
			} else {
				s.xWins = true
			}
		}
	case O:
		counts.oInCol++
		s.oCount++
		if counts.oInCol > 2 {
			if s.xWins {
				return false
			} else {
				s.oWins = true
			}
		}
	}
	return true
}

func main() {
	fmt.Print("Enter board state:")
	g := Grid{}
	text := stringFromConsole(os.Stdin)
	g.FromString(text)
	g.Print(os.Stdout)
	fmt.Print("Enter your move: ")
	text = stringFromConsole(os.Stdin)
	g.HumanMove(text, X)
	g.Print(os.Stdout)
}

func stringFromConsole(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	return scanner.Text()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
