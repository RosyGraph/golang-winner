package sequence

import "testing"

// 3, 0, -3, -6, -9, -12, -15, -18, -21, -24
func TestNthArithmetic(t *testing.T) {
	tc := []struct {
		n    int
		n1   int
		d    int
		want int
	}{
		{n: 1, n1: 1, d: 1, want: 1},
		{n: 2, n1: 0, d: 1, want: 1},
		{n: 5, n1: 3, d: 2, want: 11},
		{n: 10, n1: 3, d: -3, want: -24},
	}

	for _, c := range tc {
		t.Run("NthArithmetic", func(t *testing.T) {
			got := NthArithmetic(c.n, c.n1, c.d)

			if got != c.want {
				t.Errorf("NthArithmetic(%d, %d, %d) got %d want %d",
					c.n, c.n1, c.d, got, c.want)
			}
		})
	}
}

// 2, 4
func TestNthGeometric(t *testing.T) {
	got := NthGeometric(2, 2, 2)
	want := 4

	if got != want {
		t.Errorf("NthGeometric(%d, %d, %d) got %d want %d",
			2, 2, 2, got, want)
	}
}
