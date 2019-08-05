package main

import (
	"fmt"
	"io"
	"os"
)

type Grid [3][3]string

func (g Grid) PrintState(writer io.Writer) {
	for _, row := range g {
		for _, v := range row {
			fmt.Fprintf(writer, "%v ", v)
		}
		fmt.Fprintln(writer)
	}
}

func main() {
	grid := Grid{
		{"X", "O", "X"},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	grid.PrintState(os.Stdout)
}
