package main

import (
	"fmt"
	"io"
)

type Graph struct {
	xMax int
	xMin int
	yMax int
	yMin int
}

func (g Graph) Print(w io.Writer) {
	s := `     |     
     |     
     |     
     |     
     |     
-----------
     |     
     |     
     |     
     |     
     |     `
	fmt.Fprint(w, s)
}
