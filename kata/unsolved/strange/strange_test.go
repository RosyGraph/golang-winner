package kata

import "testing"

func TestStrangeGrid(t *testing.T) {
	tc := []struct {
		name string
		row  int32
		col  int32
		want int32
	}{
		{name: "1 2", row: 1, col: 2, want: 2},
		{name: "6 3", row: 6, col: 3, want: 25},
		{name: "1 1", row: 1, col: 1, want: 0},
		{name: "2 2", row: 2, col: 2, want: 3},
		{name: "HR test case 7",
			row: 1000000011, col: 5, want: 5e7 + 58},
		{name: "HR test case 4",
			row: 2000, col: 4, want: 9997},
		// {name: "HR test case 8",
		// row: 2e9, col: 2, want: 0x2540BE3F9},
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
