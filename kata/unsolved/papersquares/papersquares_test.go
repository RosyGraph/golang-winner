package main

import "testing"

func TestSolve(t *testing.T) {
	tc := []struct {
		name string
		n, m int32
		want uint64
	}{
		{name: "3 1", n: 3, m: 1, want: 2},
		{name: "3 2", n: 3, m: 2, want: 5},
		{name: "2 3", n: 2, m: 3, want: 5},
		{name: "6 6", n: 6, m: 6, want: 35},
		{name: "12 12", n: 12, m: 12, want: 143},
		{name: "316925480 472908316", n: 316925480, m: 472908316, want: 149876695044291679},
		{name: "689715240 759842301", n: 689715240, m: 759842301, want: 524074814996367239},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := solve(c.n, c.m)
			if got != c.want {
				t.Errorf("given (%s) got %d want %d", c.name, got, c.want)
			}
		})
	}
}
