package kata

import "testing"

func TestReverseGame(t *testing.T) {
	tc := []struct {
		name string
		n, k int32
		want int32
	}{
		{name: "3 1", n: 3, k: 1, want: 2},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := ReverseGame(c.n, c.k)
			if got != c.want {
				t.Errorf("give %s: got %d want %d", c.name, got, c.want)
			}
		})
	}
}
