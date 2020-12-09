package main

import "testing"

func TestBinary(t *testing.T) {
	tc := []struct {
		name string
		n    int
		want int
	}{
		{name: "sample 1", n: 5, want: 1},
		{name: "sample 2", n: 13, want: 2},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := Binary(c.n)
			if got != c.want {
				t.Errorf("Binary(%d) got '%d' want '%d'", 5, got, c.want)
			}
		})
	}
}
