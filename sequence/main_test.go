package main

import "testing"

// 1 + 1 + 1 + 1
func TestNthArithmetic(t *testing.T) {
	tc := []struct {
		n1   int
		d    int
		want int
	}{
		{n1: 1, d: 1, want: 1},
	}

	for _, c := range tc {
		t.Run("NthArithmetic", func(t *testing.T) {
			got := NthArithmetic(c.n1, c.d)

			if got != c.want {
				t.Errorf("NthArithmetic(%d, %d) got %d want %d",
					c.n1, c.d, got, c.want)
			}
		})
	}
}
