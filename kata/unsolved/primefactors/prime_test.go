package prime

import "testing"

type TestCase struct {
	name string
	n    int64
	want int32
}

func TestPrimeCount(t *testing.T) {
	tc := []TestCase{
		{name: "1", n: 1, want: 0},
		{name: "2", n: 2, want: 1},
	}

	check := func(t *testing.T, c TestCase) {
		got := primeCount(c.n)
		if got != c.want {
			t.Errorf("got %d want %d", got, c.want)
		}
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			check(t, c)
		})
	}
}
