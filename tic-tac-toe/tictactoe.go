package main

import (
	"fmt"
	"io"
	"strings"
)

type Grid [3][3]string

func (g Grid) PrintState(writer io.Writer) {
	fmt.Fprintln(writer, "---------")
	for _, row := range g {
		fmt.Fprintf(writer, "| %s %s %s |\n", row[0], row[1], row[2])
	}
	fmt.Fprintln(writer, "---------")
}

func (g *Grid) FromString(cells string) {
	cells = strings.Trim(cells, "\"")
	for i, char := range cells {
		g[i/3][i%3] = string(char)
	}
}
