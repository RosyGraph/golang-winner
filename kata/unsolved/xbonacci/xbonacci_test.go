package xbonacci

import "testing"

func TestXbonacci(t *testing.T) {
	tc := []struct {
		name string
		n    int
		i    int
		want int
	}{
		{name: "fibonacci 1st index", n: 2, i: 1, want: 1},
		{name: "fibonacci 3rd index", n: 2, i: 3, want: 2},
		{name: "fibonacci 5th index", n: 2, i: 5, want: 5},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := Xbonacci(c.n, c.i)

			if got != c.want {
				t.Errorf("%s: got %d want %d", c.name, got, c.want)
			}
		})
	}
}
