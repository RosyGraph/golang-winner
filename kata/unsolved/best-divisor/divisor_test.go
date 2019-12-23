package main

import "testing"

// func TestBestDivisor(t *testing.T) {
// tc := []struct {
// name string
// n    int32
// want int32
// }{
// {name: "12", n: 12, want: 6},
// {name: "16", n: 16, want: 8},
// {name: "17", n: 17, want: 17},
// }

// for _, c := range tc {
// t.Run(c.name, func(t *testing.T) {
// got := bestDivisor(c.n)

// if got != c.want {
// t.Errorf("given %s: got %d want %d", c.name, got, c.want)
// }
// })
// }
// }

func TestDigitSum(t *testing.T) {
	tc := []struct {
		name string
		n    int32
		want int32
	}{
		{name: "100", n: 100, want: 1},
		{name: "19", n: 19, want: 10},
		{name: "2", n: 2, want: 2},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := digitSum(c.n)
			if got != c.want {
				t.Errorf("given %s: got %d want %d", c.name, got, c.want)
			}
		})
	}
}
