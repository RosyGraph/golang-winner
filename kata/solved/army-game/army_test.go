package army

import "testing"

func TestMinDrops(t *testing.T) {
	tc := []struct {
		name string
		n, m int32
		want int32
	}{
		{name: "2x2", n: 2, m: 2, want: 1},
		{name: "4x2", n: 4, m: 2, want: 2},
		{name: "3x2", n: 3, m: 2, want: 2},
		{name: "3x3", n: 3, m: 3, want: 4},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := gameWithCells(c.n, c.m)

			if got != c.want {
				t.Errorf("got %d want %d", got, c.want)
			}
		})
	}
}
