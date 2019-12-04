package main

import "testing"

func TestLowestTriangle(t *testing.T) {
	tc := []struct {
		name string
		b, a int
		want int
	}{
		{name: "2 2", b: 2, a: 2, want: 2},
		{name: "2 3", b: 2, a: 3, want: 3},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := lowestTriangle(c.b, c.a)

			if got != c.want {
				t.Errorf("got %d want %d", got, c.want)
			}
		})
	}
}
