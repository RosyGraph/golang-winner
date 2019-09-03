package main

import (
	"bytes"
	"testing"
)

func TestDraw(t *testing.T) {
	buffer := bytes.Buffer{}
	graph := Graph{5, 5, -5, -5}
	graph.Print(&buffer)
	got := buffer.String()
	want := `     |     
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

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
