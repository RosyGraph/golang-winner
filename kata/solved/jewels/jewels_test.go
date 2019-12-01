package kata

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	tc := []struct {
		name string
		j    string
		s    string
		want int
	}{
		{name: "example one", j: "aA", s: "aAAbbbb", want: 3},
		{name: "example two", j: "j", s: "JJ", want: 0},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := numJewelsInStones(c.j, c.s)

			if got != c.want {
				t.Errorf("got %d want %d", got, c.want)
			}
		})
	}
}
