package main

import "testing"

func TestSolve(t *testing.T) {
	tc := []struct {
		name string
		n, m int32
		want int64
	}{
		{name: "3 1", n: 3, m: 1, want: 2},
		{name: "3 2", n: 3, m: 2, want: 5},
		{name: "12 12", n: 12, m: 12, want: 143},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := solve(c.n, c.m)
			if got != c.want {
				t.Errorf("given %s got %d want %d", c.name, got, c.want)
			}
		})
	}
}
