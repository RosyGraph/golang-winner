package main

import (
	"bytes"
	"fmt"
	"os"
)

type Grid [3][3]string

func (g Grid) PrintState(b *bytes.Buffer) {
	buffer := 
	for _, row := range g {
		for _, v := range row {
			fmt.Fprintf(b, "%q ", v)
		}
		fmt.Fprintln(b)
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
