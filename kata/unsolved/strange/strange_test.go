package kata

import "testing"

func TestStrangeGrid(t *testing.T) {
	tc := []struct {
		name string
		row  int
		col  int
		want int
	}{
		{name: "6 3", row: 6, col: 3, want: 25},
		{name: "1 1", row: 1, col: 1, want: 0},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := strangeGrid(c.row, c.col)

			if got != c.want {
				t.Errorf("got %d want %d", got, c.want)
			}
		})
	}

}
