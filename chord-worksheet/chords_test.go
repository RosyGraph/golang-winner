package main

import "testing"

func TestValueOf(t *testing.T) {
	tc := []struct {
		name string
		note string
		want int
	}{
		{name: "value of A", note: "A", want: 0},
		{name: "value of C", note: "C", want: 3},
		{name: "value of C♯", note: "C♯", want: 4},
		{name: "value of D𝄪", note: "D𝄪", want: 7},
		{name: "value of B𝄫", note: "B𝄫", want: 4},
		{name: "value of E", note: "E", want: 7},
		{name: "value of F♭", note: "F♭", want: 7},
		{name: "value of F♭", note: "F♭", want: 7},
		{name: "value of G𝄪", note: "G𝄪", want: 0},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := ValueOf(c.note)

			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}
