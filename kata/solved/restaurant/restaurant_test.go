package restaurant

import "testing"

func TestRestaurant(t *testing.T) {
	tc := []struct {
		name string
		n, m int32
		want int32
	}{
		{name: "2 2", n: 2, m: 2, want: 1},
		{name: "6 9", n: 6, m: 9, want: 6},
		{name: "38 751", n: 38, m: 751, want: 28538},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			var got int32 = restaurant(c.n, c.m)
			if got != c.want {
				t.Errorf("given %s: got %d want %d", c.name, got, c.want)
			}
		})
	}
}
