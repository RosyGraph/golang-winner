package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const X = "X"
const O = "O"

var ErrInvalidInput = errors.New("Invalid input. Please enter coordinates between 1 and 3.")
var ErrCellOccupied = errors.New("Cell is occupied. Please choose another.")

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

type Randomizer interface {
	RandomChoice(arr [][2]int) [2]int
}

type DefaultRandomizer struct{}

func main() {
	var random *DefaultRandomizer
	PlayGame(os.Stdin, random)
}

func PlayGame(reader io.Reader, randomizer Randomizer) string {
	g := EmptyGrid()
	for g.GameState() == "Game not finished" {
		g.Print(os.Stdout)
		fmt.Print("Enter your move")
		fmt.Println()
		for err := g.HumanMove(reader, X); err != nil; {
			fmt.Println(err)
			g.Print(os.Stdout)
			fmt.Print("Enter your move: ")
			fmt.Println()
			err = g.HumanMove(reader, X)
		}
		if g.GameState() != "Game not finished" {
			g.Print(os.Stdout)
			fmt.Println("You win!")
			break
		}
		fmt.Println("Making easy move...")
		g.EasyMove(randomizer, O)
	}
	return g.GameState()
}

func (r *DefaultRandomizer) RandomChoice(arr [][2]int) [2]int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	i := random.Intn(len(arr))

	return arr[i]
}

func (g *Grid) HumanMove(reader io.Reader, team string) error {
	move := Scanln(reader)
	coords := strings.Split(move, " ")
	fmt.Println(move)
	row, _ := strconv.Atoi(coords[1])
	col, _ := strconv.Atoi(coords[0])
	row = 3 - row
	col -= 1
	if row > 2 || row < 0 || col > 2 || col < 0 {
		return ErrInvalidInput
	}
	if g[row][col] == X || g[row][col] == O {
		return ErrCellOccupied
	}
	g[row][col] = team
	return nil
}

func (g *Grid) EasyMove(randomizer Randomizer, team string) {
	emptyCells := [][2]int{}
	for rowIndex, row := range g {
		for colIndex, value := range row {
			if value == " " {
				emptyCell := [2]int{rowIndex, colIndex}
				emptyCells = append(emptyCells, emptyCell)
			}
		}
	}
	move := randomizer.RandomChoice(emptyCells)
	randRow := move[0]
	randCol := move[1]
	g[randRow][randCol] = team
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

func EmptyGrid() Grid {
	g := Grid{}
	g.FromString("         ")
	return g
}

func Scanln(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Scan()
	return sc.Text()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
